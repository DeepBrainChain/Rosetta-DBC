package cmd

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"rosetta-dbc/configuration"
	"rosetta-dbc/dbc"
	"rosetta-dbc/services"

	"github.com/spf13/cobra"

	"github.com/coinbase/rosetta-sdk-go/asserter"
	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
	"golang.org/x/sync/errgroup"
)

const (
	// readTimeout is the maximum duration for reading the entire
	// request, including the body.
	readTimeout = 5 * time.Second

	// writeTimeout is the maximum duration before timing out
	// writes of the response. It is reset whenever a new
	// request's header is read.
	writeTimeout = 120 * time.Second

	// idleTimeout is the maximum amount of time to wait for the
	// next request when keep-alives are enabled.
	idleTimeout = 30 * time.Second
)

var (
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run rosetta-dbc",
		RunE:  runRunCmd,
	}
)

func runRunCmd(cmd *cobra.Command, args []string) error {
	cfg, err := configuration.LoadConfiguration()
	if err != nil {
		return fmt.Errorf("%w: unable to load configuration", err)
	}

	// The asserter automatically rejects incorrectly formatted
	// requests.
	asserter, err := asserter.NewServer(
		dbc.OperationTypes,
		dbc.HistoricalBalanceSupported,
		[]*types.NetworkIdentifier{
			{
				Blockchain:           "Substrate",
				Network:              "DBC Mainnet",
				SubNetworkIdentifier: nil,
			},
			{
				Blockchain:           "Substrate",
				Network:              "DBC Testnet",
				SubNetworkIdentifier: nil,
			},
		},
		dbc.CallMethods,
		dbc.IncludeMempoolCoins,
		"",
	)
	if err != nil {
		return fmt.Errorf("%w: could not initialize server asserter", err)
	}

	// Start required services
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go handleSignals([]context.CancelFunc{cancel})

	g, ctx := errgroup.WithContext(ctx)

	var api *dbc.API
	api, err = dbc.NewClient(cfg.DBCURL)
	if err != nil {
		return fmt.Errorf("%w: cannot initialize dbc client", err)
	}

	// 注册router
	router := services.NewBlockchainRouter(cfg, api, asserter)

	loggedRouter := server.LoggerMiddleware(router)
	loggedRouter2 := LoggerMiddleware2(loggedRouter)
	corsRouter := server.CorsMiddleware(loggedRouter2)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      corsRouter,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}

	g.Go(func() error {
		log.Printf("server listening on port %d", cfg.Port)
		return server.ListenAndServe()
	})

	g.Go(func() error {
		// If we don't shutdown server in errgroup, it will
		// never stop because server.ListenAndServe doesn't
		// take any context.
		<-ctx.Done()

		return server.Shutdown(ctx)
	})

	err = g.Wait()
	if SignalReceived {
		return errors.New("rosetta-dbc halted")
	}

	return err
}

// LoggerMiddleware is a simple logger middleware that prints the requests in
// an ad-hoc fashion to the stdlib's log.
func LoggerMiddleware2(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := readBody(r)
		inner.ServeHTTP(w, r)
		log.Printf("Got post body: %s", body)
	})
}

func readBody(req *http.Request) string {
	bodyBytes, _ := ioutil.ReadAll(req.Body)
	// Restore the io.ReadCloser to its original state
	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	// Use the content
	return string(bodyBytes)
}
