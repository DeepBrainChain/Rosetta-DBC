package configuration

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/coinbase/rosetta-sdk-go/types"

	"rosetta-dbc/dbc"
)

// Mode is the setting that determines if
// the implementation is "online" or "offline".
type Mode string

const (
	// Online is when the implementation is permitted
	// to make outbound connections.
	Online Mode = "ONLINE"

	// Offline is when the implementation is not permitted
	// to make outbound connections.
	Offline Mode = "OFFLINE"

	// Mainnet is the Ethereum Mainnet.
	Mainnet string = "MAINNET"

	// Ropsten is the Ethereum Ropsten testnet.
	Ropsten string = "ROPSTEN"

	// Rinkeby is the Ethereum Rinkeby testnet.
	Rinkeby string = "RINKEBY"

	// Goerli is the Ethereum Görli testnet.
	Goerli string = "GOERLI"

	// Testnet defaults to `Ropsten` for backwards compatibility.
	Testnet string = "TESTNET"

	// DataDirectory is the default location for all
	// persistent data.
	DataDirectory = "/data"

	// ModeEnv is the environment variable read
	// to determine mode.
	ModeEnv = "MODE"

	// NetworkEnv is the environment variable
	// read to determine network.
	NetworkEnv = "NETWORK"

	// PortEnv is the environment variable
	// read to determine the port for the Rosetta
	// implementation.
	PortEnv = "PORT"

	// GethEnv is an optional environment variable
	// used to connect rosetta-ethereum to an already
	// running geth node.
	GethEnv = "GETH"

	// TODO: add local url
	// DefaultGethURL is the default URL for
	// a running geth node. This is used
	// when GethEnv is not populated.
	// DefaultDBCURL = "http://localhost:8545"
	DefaultDBCURL = "wss://info.dbcwallet.io"

	// SkipGethAdminEnv is an optional environment variable
	// to skip geth `admin` calls which are typically not supported
	// by hosted node services. When not set, defaults to false.
	SkipGethAdminEnv = "SKIP_GETH_ADMIN"

	// MiddlewareVersion is the version of rosetta-ethereum.
	MiddlewareVersion = "0.0.4"
)

type Configuration struct {
	Mode                   Mode
	Network                *types.NetworkIdentifier
	GenesisBlockIdentifier *types.BlockIdentifier
	DBCURL                 string
	RemoteGeth             bool
	Port                   int
	GethArguments          string
	SkipGethAdmin          bool
	GethHeaders            []*dbc.HTTPHeader

	// Block Reward Data
	// Params *params.ChainConfig
	Params string
}

func LoadConfiguration() (*Configuration, error) {
	config := &Configuration{}

	modeValue := Mode(os.Getenv(ModeEnv))
	switch modeValue {
	case Online:
		config.Mode = Online
	case Offline:
		config.Mode = Offline
	case "":
		return nil, errors.New("MODE must be populated")
	default:
		return nil, fmt.Errorf("%s is not a valid mode", modeValue)
	}

	networkValue := os.Getenv(NetworkEnv)
	switch networkValue {
	case Mainnet:
		config.Network = &types.NetworkIdentifier{
			Blockchain: dbc.Blockchain,
			Network:    dbc.MainnetNetwork,
		}
		config.GenesisBlockIdentifier = dbc.MainnetGenesisBlockIdentifier
		config.Params = dbc.MainnetChainConfig
		config.GethArguments = dbc.MainnetArguments
	case "":
		return nil, errors.New("NETWORK must be populated")
	default:
		return nil, fmt.Errorf("%s is not a valid network", networkValue)
	}

	config.DBCURL = DefaultDBCURL

	portValue := os.Getenv(PortEnv)
	if len(portValue) == 0 {
		return nil, errors.New("PORT must be populated")
	}

	port, err := strconv.Atoi(portValue)
	if err != nil || len(portValue) == 0 || port <= 0 {
		return nil, fmt.Errorf("%w: unable to parse port %s", err, portValue)
	}
	config.Port = port

	return config, nil
}
