package services

import (
	"context"

	"rosetta-dbc/configuration"
	"rosetta-dbc/dbc"

	// "github.com/centrifuge/go-substrate-rpc-client/v3/client"
	// "github.com/centrifuge/go-substrate-rpc-client/v3/rpc"

	"github.com/coinbase/rosetta-sdk-go/asserter"
	"github.com/coinbase/rosetta-sdk-go/types"
)

type NetworkAPIService struct {
	config *configuration.Configuration
	client Client
}

func NewNetworkAPIService(cfg *configuration.Configuration, client Client) *NetworkAPIService {
	return &NetworkAPIService{
		config: cfg,
		client: client,
	}
}

// NetworkList implements the /network/list endpoint
func (s *NetworkAPIService) NetworkList(
	ctx context.Context,
	request *types.MetadataRequest,
) (*types.NetworkListResponse, *types.Error) {
	return &types.NetworkListResponse{
		NetworkIdentifiers: []*types.NetworkIdentifier{s.config.Network},
	}, nil
}

// NetworkOptions implements the /network/options endpoint.
func (s *NetworkAPIService) NetworkOptions(
	ctx context.Context,
	request *types.NetworkRequest,
) (*types.NetworkOptionsResponse, *types.Error) {
	return &types.NetworkOptionsResponse{
		Version: &types.Version{
			NodeVersion:       dbc.NodeVersion,
			RosettaVersion:    types.RosettaAPIVersion,
			MiddlewareVersion: types.String(configuration.MiddlewareVersion),
		},
		Allow: &types.Allow{
			Errors:                  Errors,
			OperationTypes:          dbc.OperationTypes,
			OperationStatuses:       dbc.OperationStatuses,
			HistoricalBalanceLookup: dbc.HistoricalBalanceSupported,
			CallMethods:             dbc.CallMethods,
		},
	}, nil
}

// NetworkStatus implements the /network/status endpoint.
func (s *NetworkAPIService) NetworkStatus(
	ctx context.Context,
	request *types.NetworkRequest,
) (*types.NetworkStatusResponse, *types.Error) {
	if s.config.Mode != configuration.Online {
		return nil, ErrGethNotReady
	}

	currentBlock, currentTime, syncStatus, peers, err := s.client.Status(ctx)
	if err != nil {
		return nil, wrapErr(ErrGeth, err)
	}

	if currentTime < asserter.MinUnixEpoch {
		return nil, ErrGethNotReady
	}

	return &types.NetworkStatusResponse{
		CurrentBlockIdentifier: currentBlock,
		CurrentBlockTimestamp:  currentTime,
		GenesisBlockIdentifier: s.config.GenesisBlockIdentifier,
		SyncStatus:             syncStatus,
		Peers:                  peers,
	}, nil
}
