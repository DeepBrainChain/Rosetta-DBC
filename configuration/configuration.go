package configuration

import (
	"github.com/coinbase/rosetta-sdk-go/types"
	"rosetta-dbc/dbc"
)

const (
	// MiddlewareVersion is the version of rosetta-ethereum.
	MiddlewareVersion = "0.1"
)

// Mode is the setting that determines if
// the implementation is "online" or "offline".
type Mode string

type Configuration struct {
	Mode                   Mode
	Network                *types.NetworkIdentifier
	GenesisBlockIdentifier *types.BlockIdentifier
	GethURL                string
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
	return config, nil
}
