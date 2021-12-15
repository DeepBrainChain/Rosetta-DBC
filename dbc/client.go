package dbc

import (
	"context"
	"math/big"
	"time"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v3"
	gsTypes "github.com/centrifuge/go-substrate-rpc-client/v3/types"
	RosettaTypes "github.com/coinbase/rosetta-sdk-go/types"

	// "github.com/vedhavyas/go-subkey"
	// "github.com/vedhavyas/go-subkey/sr25519"
	subscanSS58 "github.com/itering/subscan/util/ss58"
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
	if err != nil {
		return nil, -1, nil, nil, err
	}

	timestamp, err := ec.getBlockTimestamp(uint64(block.Block.Header.Number))
	if err != nil {
		return nil, -1, nil, nil, err
	}

	peers := []*RosettaTypes.Peer{}

	allPeer, err := ec.RPC.System.Peers()
	if err != nil {
		return nil, -1, nil, nil, err
	}
	for i := 0; i < len(allPeer); i++ {
		peers = append(peers, &RosettaTypes.Peer{
			PeerID:   string(allPeer[i].PeerID),
			Metadata: map[string]interface{}{},
		})
	}

	syncStatus := &RosettaTypes.SyncStatus{
		// TODO: add this
		// CurrentIndex *int64 `json:"current_index,omitempty"`
		// 	TargetIndex *int64 `json:"target_index,omitempty"`
		// 	Stage *string `json:"stage,omitempty"`
		Synced: &health.IsSyncing,
	}

	return &RosettaTypes.BlockIdentifier{
			Hash:  blockHash.Hex(),
			Index: int64(block.Block.Header.Number),
		},
		timestamp,
		syncStatus,
		peers,
		nil

}

func ss58ToPubkey(ss58Addr string) ([]byte, error) {
	pubkey := subscanSS58.Decode(ss58Addr, 42)
	return gsTypes.HexDecodeString(pubkey)
}

func (ec *API) latestBlockIdentifier() (*RosettaTypes.BlockIdentifier, error) {
	block, err := ec.RPC.Chain.GetBlockLatest()
	if err != nil {
		return nil, err
	}

	blockHash, err := ec.RPC.Chain.GetBlockHash(uint64(block.Block.Header.Number))
	if err != nil {
		return nil, err
	}

	return &RosettaTypes.BlockIdentifier{
		Hash:  blockHash.Hex(),
		Index: int64(block.Block.Header.Number),
	}, nil
}

func (ec *API) latestMeta() (*gsTypes.Metadata, error) {
	meta, err := ec.RPC.State.GetMetadataLatest()
	return meta, err
}

func (ec *API) getBlockTimestamp(blockHeight uint64) (int64, error) {
	meta, err := ec.latestMeta()
	if err != nil {
		return 0, err
	}

	key, err := gsTypes.CreateStorageKey(meta, "Timestamp", "Now")

	blockHash, err := ec.RPC.Chain.GetBlockHash(blockHeight)
	if err != nil {
		return 0, err
	}

	var timestamp gsTypes.Moment
	ok, err := ec.RPC.State.GetStorage(key, &timestamp, blockHash)
	if err != nil || !ok {
		return 0, err
	}

	return timestamp.UnixMilli(), nil
}

// TODO: add get block transaction
func (ec *API) getBlockTransactions() ([]*RosettaTypes.Transaction, error) {
	return nil, nil
}

func (ec *API) Balance(
	ctx context.Context,
	account *RosettaTypes.AccountIdentifier,
	block *RosettaTypes.PartialBlockIdentifier,
) (*RosettaTypes.AccountBalanceResponse, error) {
	meta, err := ec.latestMeta()
	if err != nil {
		return nil, err
	}

	pubkey, err := ss58ToPubkey(account.Address)
	if err != nil {
		return nil, err
	}

	key, err := gsTypes.CreateStorageKey(meta, "System", "Account", pubkey)
	if err != nil {
		return nil, err
	}

	var accountInfo gsTypes.AccountInfo
	ok, err := ec.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil || !ok {
		return nil, err
	}

	blockIdentifier, err := ec.latestBlockIdentifier()
	if err != nil {
		return nil, err
	}

	return &RosettaTypes.AccountBalanceResponse{
		BlockIdentifier: blockIdentifier,
		Balances: []*RosettaTypes.Amount{
			{
				Value:    accountInfo.Data.Free.String(),
				Currency: Currency,
			},
		},
		Metadata: nil,
	}, nil
}

func (ec *API) Block(
	ctx context.Context,
	blockIdentifier *RosettaTypes.PartialBlockIdentifier,
) (*RosettaTypes.Block, error) {
	var parentBlockIdentifier *RosettaTypes.BlockIdentifier

	if *blockIdentifier.Index == 0 {
		parentBlockIdentifier = MainnetGenesisBlockIdentifier
	} else {
		parentBlockHash, err := ec.RPC.Chain.GetBlockHash(uint64(*blockIdentifier.Index) - 1)
		if err != nil {
			return nil, err
		}

		parentBlockIdentifier = &RosettaTypes.BlockIdentifier{
			Index: *blockIdentifier.Index - 1,
			Hash:  parentBlockHash.Hex(),
		}
	}

	blockHash, err := ec.RPC.Chain.GetBlockHash(uint64(*blockIdentifier.Index))
	if err != nil {
		return nil, err
	}

	timestamp, err := ec.getBlockTimestamp(uint64(*blockIdentifier.Index))
	if err != nil {
		return nil, err
	}

	return &RosettaTypes.Block{
		BlockIdentifier: &RosettaTypes.BlockIdentifier{
			Index: *blockIdentifier.Index,
			Hash:  blockHash.Hex(),
		},
		ParentBlockIdentifier: parentBlockIdentifier,
		Timestamp:             timestamp,
		Transactions:          []*RosettaTypes.Transaction{}, // TODO:
		Metadata:              nil,                           // TODO:
	}, nil
}

func (ec *API) Call(
	ctx context.Context,
	request *RosettaTypes.CallRequest,
) (*RosettaTypes.CallResponse, error) {
	// 1. if is query, return query result

	// 2. if is extrinsic, sign and submit
	// ec.RPC.Author.SubmitExtrinsic(xt gsTypes.Extrinsic)

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
