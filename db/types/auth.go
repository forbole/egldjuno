package types

import "reflect"

// AccountRow represents a single row of the account table
type AccountRow struct {
	Address                  string `db:"address"`
	Balance                  string `db:"balance"`
	Nonce                    string `db:"nonce"`
	Shard                    int    `db:"shard"`
	ScamInfo                 []byte `db:"scam_info"`
	Code                     string `db:"code"`
	CodeHash                 string `db:"code_hash"`
	RootHash                 string `db:"root_hash"`
	TxCount                  int    `db:"tx_count"`
	ScrCount                 int    `db:"scr_count"`
	Username                 string `db:"username"`
	DeveloperReward          string `db:"developer_reward"`
	OwnerAddress             string `db:"owner_address"`
	DeployedAt               int    `db:"deployed_at"`
	IsUpgradeable            bool   `db:"is_upgradeable"`
	IsReadable               bool   `db:"is_readable"`
	IsPayable                bool   `db:"is_payable"`
	IsPayableBySmartContract bool   `db:"is_payable_by_smart_contract"`
}

// Equal tells whether v and w represent the same rows
func (v AccountRow) Equal(w AccountRow) bool {
	return v.Address == w.Address &&
		v.Balance == w.Balance &&
		v.Nonce == w.Nonce &&
		v.Shard == w.Shard &&
		reflect.DeepEqual(v.ScamInfo, w.ScamInfo) &&
		v.Code == w.Code &&
		v.CodeHash == w.CodeHash &&
		v.RootHash == w.RootHash &&
		v.TxCount == w.TxCount &&
		v.ScrCount == w.ScrCount &&
		v.Username == w.Username &&
		v.DeveloperReward == w.DeveloperReward &&
		v.OwnerAddress == w.OwnerAddress &&
		v.DeployedAt == w.DeployedAt &&
		v.IsUpgradeable == w.IsUpgradeable &&
		v.IsReadable == w.IsReadable &&
		v.IsPayable == w.IsPayable &&
		v.IsPayableBySmartContract == w.IsPayableBySmartContract
}

// AccountRow allows to build a new AccountRow
func NewAccountRow(
	address string,
	balance string,
	nonce string,
	shard int,
	scamInfo []byte,
	code string,
	codeHash string,
	rootHash string,
	txCount int,
	scrCount int,
	username string,
	developerReward string,
	ownerAddress string,
	deployedAt int,
	isUpgradeable bool,
	isReadable bool,
	isPayable bool,
	isPayableBySmartContract bool) AccountRow {
	return AccountRow{
		Address:                  address,
		Balance:                  balance,
		Nonce:                    nonce,
		Shard:                    shard,
		ScamInfo:                 scamInfo,
		Code:                     code,
		CodeHash:                 codeHash,
		RootHash:                 rootHash,
		TxCount:                  txCount,
		ScrCount:                 scrCount,
		Username:                 username,
		DeveloperReward:          developerReward,
		OwnerAddress:             ownerAddress,
		DeployedAt:               deployedAt,
		IsUpgradeable:            isUpgradeable,
		IsReadable:               isReadable,
		IsPayable:                isPayable,
		IsPayableBySmartContract: isPayableBySmartContract,
	}
}

// BlockRow represents a single row of the block table
type BlockRow struct {
	Hash          string `db:"hash"`
	Epoch         int64  `db:"epoch"`
	Nonce         int64  `db:"nonce"`
	PrevHash      string `db:"prev_hash"`
	Proposer      string `db:"proposer"`
	PubKeyBitmap  string `db:"pub_key_bitmap"`
	Round         int64  `db:"round"`
	Shard         int64  `db:"shard"`
	Size          int64  `db:"size"`
	SizeTxs       int64  `db:"size_txs"`
	StateRootHash string `db:"state_root_hash"`
	TimeStamp     int64  `db:"time_stamp"`
	TxCount       int64  `db:"tx_count"`
	GasConsumed   int64  `db:"gas_consumed"`
	GasRefunded   int64  `db:"gas_refunded"`
	GasPenalized  int64  `db:"gas_penalized"`
	MaxGasLimit   int64  `db:"max_gas_limit"`
}

// Equal tells whether v and w represent the same rows
func (v BlockRow) Equal(w BlockRow) bool {
	return v.Hash == w.Hash &&
		v.Epoch == w.Epoch &&
		v.Nonce == w.Nonce &&
		v.PrevHash == w.PrevHash &&
		v.Proposer == w.Proposer &&
		v.PubKeyBitmap == w.PubKeyBitmap &&
		v.Round == w.Round &&
		v.Shard == w.Shard &&
		v.Size == w.Size &&
		v.SizeTxs == w.SizeTxs &&
		v.StateRootHash == w.StateRootHash &&
		v.TimeStamp == w.TimeStamp &&
		v.TxCount == w.TxCount &&
		v.GasConsumed == w.GasConsumed &&
		v.GasRefunded == w.GasRefunded &&
		v.GasPenalized == w.GasPenalized &&
		v.MaxGasLimit == w.MaxGasLimit
}

// BlockRow allows to build a new BlockRow
func NewBlockRow(
	hash string,
	epoch int64,
	nonce int64,
	prevHash string,
	proposer string,
	pubKeyBitmap string,
	round int64,
	shard int64,
	size int64,
	sizeTxs int64,
	stateRootHash string,
	timeStamp int64,
	txCount int64,
	gasConsumed int64,
	gasRefunded int64,
	gasPenalized int64,
	maxGasLimit int64) BlockRow {
	return BlockRow{
		Hash:          hash,
		Epoch:         epoch,
		Nonce:         nonce,
		PrevHash:      prevHash,
		Proposer:      proposer,
		PubKeyBitmap:  pubKeyBitmap,
		Round:         round,
		Shard:         shard,
		Size:          size,
		SizeTxs:       sizeTxs,
		StateRootHash: stateRootHash,
		TimeStamp:     timeStamp,
		TxCount:       txCount,
		GasConsumed:   gasConsumed,
		GasRefunded:   gasRefunded,
		GasPenalized:  gasPenalized,
		MaxGasLimit:   maxGasLimit,
	}
}

// TokenBalanceRow represents a single row of the token_balance table
type TokenBalanceRow struct {
	Address    string `db:"address"`
	Identifier string `db:"identifier"`
	Balance    string `db:"balance"`
}

// Equal tells whether v and w represent the same rows
func (v TokenBalanceRow) Equal(w TokenBalanceRow) bool {
	return v.Address == w.Address &&
		v.Identifier == w.Identifier &&
		v.Balance == w.Balance
}

// TokenBalanceRow allows to build a new TokenBalanceRow
func NewTokenBalanceRow(
	address string,
	identifier string,
	balance string) TokenBalanceRow {
	return TokenBalanceRow{
		Address:    address,
		Identifier: identifier,
		Balance:    balance,
	}
}

// AccountNftRow represents a single row of the account_nft table
type AccountNftRow struct {
	Address    string `db:"address"`
	Identifier string `db:"identifier"`
}

// Equal tells whether v and w represent the same rows
func (v AccountNftRow) Equal(w AccountNftRow) bool {
	return v.Address == w.Address &&
		v.Identifier == w.Identifier
}

// AccountNftRow allows to build a new AccountNftRow
func NewAccountNftRow(
	address string,
	identifier string) AccountNftRow {
	return AccountNftRow{
		Address:    address,
		Identifier: identifier,
	}
}

// AccountContractRow represents a single row of the account_contract table
type AccountContractRow struct {
	Address         string `db:"address"`
	ContractAddress string `db:"contract_address"`
	DeployTxHash    string `db:"deploy_tx_hash"`
	Timestamp       int    `db:"timestamp"`
}

// Equal tells whether v and w represent the same rows
func (v AccountContractRow) Equal(w AccountContractRow) bool {
	return v.Address == w.Address &&
		v.ContractAddress == w.ContractAddress &&
		v.DeployTxHash == w.DeployTxHash &&
		v.Timestamp == w.Timestamp
}

// AccountContractRow allows to build a new AccountContractRow
func NewAccountContractRow(
	address string,
	contractAddress string,
	deployTxHash string,
	timestamp int) AccountContractRow {
	return AccountContractRow{
		Address:         address,
		ContractAddress: contractAddress,
		DeployTxHash:    deployTxHash,
		Timestamp:       timestamp,
	}
}
