package services

import (
	"net/http"

	"rosetta-dbc/configuration"

	"github.com/coinbase/rosetta-sdk-go/asserter"
	"github.com/coinbase/rosetta-sdk-go/server"
)

func NewBlockchainRouter(
	config *configuration.Configuration,
	client Client,
	asserter *asserter.Asserter,
) http.Handler {
	networkAPIService := NewNetworkAPIService(config, client)
	networkAPIController := server.NewNetworkAPIController(
		networkAPIService,
		asserter,
	)

	accountAPIService := NewAccountAPIService(config, client)
	accountAPIController := server.NewAccountAPIController(
		accountAPIService,
		asserter,
	)

	blockAPIService := NewBlockAPIService(config, client)
	blockAPIController := server.NewBlockAPIController(
		blockAPIService,
		asserter,
	)

	constructionAPIService := NewConstructionAPIService(config, client)
	constructionAPIController := server.NewConstructionAPIController(
		constructionAPIService,
		asserter,
	)

	mempoolAPIService := NewMempoolAPIService()
	mempoolAPIController := server.NewMempoolAPIController(
		mempoolAPIService,
		asserter,
	)

	callAPIService := NewCallAPIService(config, client)
	callAPIController := server.NewCallAPIController(
		callAPIService,
		asserter,
	)

	return server.NewRouter(
		networkAPIController,
		accountAPIController,
		blockAPIController,
		constructionAPIController,
		mempoolAPIController,
		callAPIController,
	)

}
