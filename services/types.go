package services

import (
	"context"
	gsTypes "github.com/centrifuge/go-substrate-rpc-client/v3/types"
	"github.com/coinbase/rosetta-sdk-go/types"
	"math/big"
)

// Client is used by the servicers to get block
// data and to submit transactions.
type Client interface {
	Status(context.Context) (
		*types.BlockIdentifier,
		int64,
		*types.SyncStatus,
		[]*types.Peer,
		error,
	)

	Block(
		context.Context,
		*types.PartialBlockIdentifier,
	) (*types.Block, error)

	Balance(
		context.Context,
		*types.AccountIdentifier,
		*types.PartialBlockIdentifier,
	) (*types.AccountBalanceResponse, error)

	PendingNonceAt(context.Context, gsTypes.Address) (uint64, error)

	SuggestGasPrice(ctx context.Context) (*big.Int, error)

	SendTransaction(ctx context.Context, tx *gsTypes.Extrinsic) error

	Call(
		ctx context.Context,
		request *types.CallRequest,
	) (*types.CallResponse, error)
}
