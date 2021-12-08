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
	// Blockchain is Ethereum.
	Blockchain string = "Substrate"

	// MainnetNetwork is the value of the network
	// in MainnetNetworkIdentifier.
	MainnetNetwork string = "DBC Mainnet"

	// TODO: 增加这个Identifier
	// TestnetNetwork is the value of the network
	// in TestnetNetworkIdentifier
	TestnetNetwork string = "DBC Testnet"

	// UnclesRewardMultiplier is the uncle reward
	// multiplier.
	UnclesRewardMultiplier = 32

	// MaxUncleDepth is the maximum depth for
	// an uncle to be rewarded.
	MaxUncleDepth = 8

	// GenesisBlockIndex is the index of the
	// genesis block.
	GenesisBlockIndex = int64(0)

	// TODO: change here
	// TransferGasLimit is the gas limit
	// of a transfer.
	TransferGasLimit = int64(21000) //nolint:gomnd

	// TODO: change here
	// MainnetGethArguments are the arguments to start a mainnet geth instance.
	MainnetArguments = ``
	TestnetArguments = ``
)

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

	// MainnetGenesisHash = rpc.types.NewHash(MustHexDecodeString("0xdcd1346701ca8396496e52aa2785b1748deb6db09551b72159dcb3e08991025b"))

	MainnetGenesisHash = "0xd523fa2e0581f069b4f0c7b5944c21e9abc72305a08067868c91b898d1bf1dff"

	TestnetGenesisHash = "0xe57c5ceeba24c5c5e525db2bcad3c661a53cd345aeaa8a2fbc66347ced360027"
)

var (
	// MainnetGenesisBlockIdentifier is the *types.BlockIdentifier
	// of the mainnet genesis block.
	MainnetGenesisBlockIdentifier = &types.BlockIdentifier{
		Hash:  MainnetGenesisHash,
		Index: GenesisBlockIndex,
	}

	TestnetGenesisBlockIdentifier = &types.BlockIdentifier{
		Hash:  TestnetGenesisHash,
		Index: GenesisBlockIndex,
	}

	MainnetChainConfig = ""
	TestnetChainConfig = ""

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

	// CallMethods are all supported call methods.
	CallMethods = []string{
		"dbc_getBlockByNumber",
		"dbc_getTransactionReceipt",
		"dbc_call",
		"dbc_estimateGas", // TODO: change here
	}
)
