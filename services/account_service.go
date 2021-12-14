package services

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"
	"rosetta-dbc/configuration"
)

// AccountAPIService implements the server.AccountAPIServicer interface.
type AccountAPIService struct {
	config *configuration.Configuration
	client Client
}

// NewAccountAPIService returns a new *AccountAPIService.
func NewAccountAPIService(
	cfg *configuration.Configuration,
	client Client,
) *AccountAPIService {
	return &AccountAPIService{
		config: cfg,
		client: client,
	}
}

// AccountBalance implements /account/balance.
func (s *AccountAPIService) AccountBalance(
	ctx context.Context,
	request *types.AccountBalanceRequest,
) (*types.AccountBalanceResponse, *types.Error) {
	balanceResponse, err := s.client.Balance(
		ctx,
		request.AccountIdentifier,
		request.BlockIdentifier,
	)
	if err != nil {
		return nil, wrapErr(ErrDBC, err)
	}

	return balanceResponse, nil
}

// AccountCoins implements /account/coins.
func (s *AccountAPIService) AccountCoins(
	ctx context.Context,
	request *types.AccountCoinsRequest,
) (*types.AccountCoinsResponse, *types.Error) {
	return nil, wrapErr(ErrUnimplemented, nil)
}
