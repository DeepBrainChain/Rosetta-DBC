package dbc

import (
	"context"
	"math/big"
	"time"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v3"
	gsTypes "github.com/centrifuge/go-substrate-rpc-client/v3/types"
	RosettaTypes "github.com/coinbase/rosetta-sdk-go/types"
)

const (
	gethHTTPTimeout = 120 * time.Second

	maxTraceConcurrency  = int64(16) // nolint:gomnd
	semaphoreTraceWeight = int64(1)  // nolint:gomnd
)

type API gsrpc.SubstrateAPI

func NewClient(url string) (*API, error) {
	api, err := gsrpc.NewSubstrateAPI(url)
	if err != nil {
		return nil, err
	}
	return &API{
		RPC:    api.RPC,
		Client: api.Client,
	}, err
}

func (ec *API) Close() {}

func (ec *API) Status(ctx context.Context) (
	*RosettaTypes.BlockIdentifier,
	int64,
	*RosettaTypes.SyncStatus,
	[]*RosettaTypes.Peer,
	error,

) {
	block, err := ec.RPC.Chain.GetBlockLatest()
	if err != nil {
		return nil, -1, nil, nil, err
	}

	blockHash, err := ec.RPC.Chain.GetBlockHash(uint64(block.Block.Header.Number))
	if err != nil {
		return nil, -1, nil, nil, err
	}

	health, err := ec.RPC.System.Health()

	// peers, _ := ec.RPC.System.Peers()
	// Dont need any peers for now, format response
	peers := []*RosettaTypes.Peer{}

	syncStatus := &RosettaTypes.SyncStatus{
		// CurrentIndex *int64 `json:"current_index,omitempty"`
		// 	TargetIndex *int64 `json:"target_index,omitempty"`
		// 	Stage *string `json:"stage,omitempty"`
		Synced: &health.IsSyncing,
	}

	return &RosettaTypes.BlockIdentifier{
			Hash:  blockHash.Hex(),
			Index: int64(block.Block.Header.Number),
		},
		1, // convertTime(header.Time),
		syncStatus,
		peers,
		nil

}

func (ec *API) Balance(
	ctx context.Context,
	account *RosettaTypes.AccountIdentifier,
	block *RosettaTypes.PartialBlockIdentifier,
) (*RosettaTypes.AccountBalanceResponse, error) {
	return nil, nil
}

func (ec *API) Block(
	ctx context.Context,
	blockIdentifier *RosettaTypes.PartialBlockIdentifier,
) (*RosettaTypes.Block, error) {
	return nil, nil
}

func (ec *API) Call(
	ctx context.Context,
	request *RosettaTypes.CallRequest,
) (*RosettaTypes.CallResponse, error) {
	return nil, nil
}

// PendingNonceAt returns the account nonce of the given account in the pending state.
// This is the nonce that should be used for the next transaction.
func (ec *API) PendingNonceAt(ctx context.Context, account gsTypes.Address) (uint64, error) {
	return 0, nil
}

func (ec *API) SendTransaction(ctx context.Context, tx *gsTypes.Extrinsic) error {
	return nil
}

func (ec *API) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return nil, nil
}
