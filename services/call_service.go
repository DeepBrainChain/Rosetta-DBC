package services

import (
	"context"
	"errors"

	"github.com/coinbase/rosetta-sdk-go/types"
	"rosetta-dbc/configuration"
	"rosetta-dbc/dbc"
)

// CallAPIService implements the server.CallAPIServicer interface.
type CallAPIService struct {
	config *configuration.Configuration
	client Client
}

// NewCallAPIService creates a new instance of a CallAPIService.
func NewCallAPIService(cfg *configuration.Configuration, client Client) *CallAPIService {
	return &CallAPIService{
		config: cfg,
		client: client,
	}
}

// Call implements the /call endpoint.
func (s *CallAPIService) Call(
	ctx context.Context,
	request *types.CallRequest,
) (*types.CallResponse, *types.Error) {
	response, err := s.client.Call(ctx, request)
	if errors.Is(err, dbc.ErrCallParametersInvalid) {
		return nil, wrapErr(ErrCallParametersInvalid, err)
	}
	if errors.Is(err, dbc.ErrCallOutputMarshal) {
		return nil, wrapErr(ErrCallOutputMarshal, err)
	}
	if errors.Is(err, dbc.ErrCallMethodInvalid) {
		return nil, wrapErr(ErrCallMethodInvalid, err)
	}
	if err != nil {
		return nil, wrapErr(ErrDBC, err)
	}

	return response, nil
}
