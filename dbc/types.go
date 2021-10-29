package dbc

import (
	"context"

	rpc "github.com/centrifuge/go-substrate-rpc-client/v3/gethrpc"
	"github.com/coinbase/rosetta-sdk-go/types"
)

// JSONRPC is the interface for accessing go-ethereum's JSON RPC endpoint.
type JSONRPC interface {
	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
	BatchCallContext(ctx context.Context, b []rpc.BatchElem) error
	Close()
}

type GraphQL interface {
	Query(ctx context.Context, input string) (string, error)
}

type HTTPHeader struct {
	Key   string
	Value string
}

const (
	// NodeVersion is the version of geth we are using.
	NodeVersion = "264"

	// Symbol is the symbol value
	// used in Currency.
	Symbol = "DBC"

	// Decimals is the decimals value
	// used in Currency.
	Decimals = 15

	// MinerRewardOpType is used to describe
	// a miner block reward.
	MinerRewardOpType = "MINER_REWARD"

	// UncleRewardOpType is used to describe
	// an uncle block reward.
	UncleRewardOpType = "UNCLE_REWARD"

	// FeeOpType is used to represent fee operations.
	FeeOpType = "FEE"

	// CallOpType is used to represent CALL trace operations.
	CallOpType = "CALL"

	// CreateOpType is used to represent CREATE trace operations.
	CreateOpType = "CREATE"

	// Create2OpType is used to represent CREATE2 trace operations.
	Create2OpType = "CREATE2"

	// SelfDestructOpType is used to represent SELFDESTRUCT trace operations.
	SelfDestructOpType = "SELFDESTRUCT"

	// CallCodeOpType is used to represent CALLCODE trace operations.
	CallCodeOpType = "CALLCODE"

	// DelegateCallOpType is used to represent DELEGATECALL trace operations.
	DelegateCallOpType = "DELEGATECALL"

	// StaticCallOpType is used to represent STATICCALL trace operations.
	StaticCallOpType = "STATICCALL"

	// DestructOpType is a synthetic operation used to represent the
	// deletion of suicided accounts that still have funds at the end
	// of a transaction.
	DestructOpType = "DESTRUCT"

	// SuccessStatus is the status of any
	// Ethereum operation considered successful.
	SuccessStatus = "SUCCESS"

	// FailureStatus is the status of any
	// Ethereum operation considered unsuccessful.
	FailureStatus = "FAILURE"

	// HistoricalBalanceSupported is whether
	// historical balance is supported.
	HistoricalBalanceSupported = true

	// IncludeMempoolCoins does not apply to rosetta-ethereum as it is not UTXO-based.
	IncludeMempoolCoins = false
)

var (

	// Currency is the *types.Currency for all
	// Ethereum networks.
	Currency = &types.Currency{
		Symbol:   Symbol,
		Decimals: Decimals,
	}

	// OperationTypes are all suppoorted operation types.
	OperationTypes = []string{
		MinerRewardOpType,
		UncleRewardOpType,
		FeeOpType,
		CallOpType,
		CreateOpType,
		Create2OpType,
		SelfDestructOpType,
		CallCodeOpType,
		DelegateCallOpType,
		StaticCallOpType,
		DestructOpType,
	}

	// OperationStatuses are all supported operation statuses.
	OperationStatuses = []*types.OperationStatus{
		{
			Status:     SuccessStatus,
			Successful: true,
		},
		{
			Status:     FailureStatus,
			Successful: false,
		},
	}

	CallMethods = []string{}
)
