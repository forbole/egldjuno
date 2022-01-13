package types

// AccountRow represents a single row of the account table
type AccountRow struct {
	Address string `db:"address"`
}

// Equal tells whether v and w represent the same rows
func (v AccountRow) Equal(w AccountRow) bool {
	return v.Address == w.Address
}

// AccountRow allows to build a new AccountRow
func NewAccountRow(
	address string) AccountRow {
	return AccountRow{
		Address: address,
	}
}

// AccountBalanceRow represents a single row inside the account table
type AccountBalanceRow struct {
	Address     string  `db:"address"`
	Balance     float64 `db:"balance"`
	Code        string  `db:"code"`
	ContractMap string  `db:"contract_map"`
	Height      uint64  `db:"height"`
}

// NewAccountBalanceRow allows to easily build a new AccountBalanceRow
func NewAccountBalanceRow(address string, balance float64, code, contractMap string, height uint64) AccountBalanceRow {
	return AccountBalanceRow{
		Address:     address,
		Balance:     balance,
		Code:        code,
		ContractMap: contractMap,
		Height:      height,
	}
}

// Equal tells whether a and b contain the same data
func (a AccountBalanceRow) Equal(b AccountBalanceRow) bool {
	return (a.Address == b.Address &&
		a.Balance == b.Balance &&
		a.Code == b.Code &&
		a.ContractMap == b.ContractMap &&
		a.Height == b.Height)
}

// LockedAccountRow represents a single row of the locked_account table
type LockedAccountRow struct {
	Address       string `db:"address"`
	LockedAddress string `db:"locked_address"`
}

// Equal tells whether v and w represent the same rows
func (v LockedAccountRow) Equal(w LockedAccountRow) bool {
	return v.Address == w.Address &&
		v.LockedAddress == w.LockedAddress
}

// LockedAccountRow allows to build a new LockedAccountRow
func NewLockedAccountRow(
	address string,
	lockedAddress string) LockedAccountRow {
	return LockedAccountRow{
		Address:       address,
		LockedAddress: lockedAddress,
	}
}

type LockedAccountBalanceRow struct {
	LockedAddress string `db:"locked_address"`
	Balance       int    `db:"balance"`
	UnlockLimit   int    `db:"unlock_limit"`
	Height        int    `db:"height"`
}

func NewLockedAccountBalanceRow(lockedAddress string, balance, unlockLimit, height int) LockedAccountBalanceRow {
	return LockedAccountBalanceRow{
		LockedAddress: lockedAddress,
		Balance:       balance,
		UnlockLimit:   unlockLimit,
		Height:        height,
	}
}

func (a LockedAccountBalanceRow) Equal(b LockedAccountBalanceRow) bool {
	return (a.Balance == b.Balance &&
		a.LockedAddress == b.LockedAddress &&
		a.UnlockLimit == b.UnlockLimit &&
		a.Height == b.Height)
}

type DelegatorAccountRow struct {
	AccountAddress  string `db:"account_address"`
	DelegatorId     int64  `db:"delegator_id"`
	DelegatorNodeId string `db:"delegator_node_id"`
}

// Equal tells whether v and w represent the same rows
func (v DelegatorAccountRow) Equal(w DelegatorAccountRow) bool {
	return v.AccountAddress == w.AccountAddress &&
		v.DelegatorId == w.DelegatorId &&
		v.DelegatorNodeId == w.DelegatorNodeId
}

// DelegatorAccountRow allows to build a new DelegatorAccountRow
func NewDelegatorAccountRow(
	accountAddress string,
	delegatorId int64,
	delegatorNodeId string) DelegatorAccountRow {
	return DelegatorAccountRow{
		AccountAddress:  accountAddress,
		DelegatorId:     delegatorId,
		DelegatorNodeId: delegatorNodeId,
	}
}

// StakerAccountRow represents a single row of the staker_account table
type StakerAccountRow struct {
	AccountAddress string `db:"account_address"`
	StakerNodeId   string `db:"staker_node_id"`
	StakerNodeInfo string `db:"staker_node_info"`
}

// Equal tells whether v and w represent the same rows
func (v StakerAccountRow) Equal(w StakerAccountRow) bool {
	return v.AccountAddress == w.AccountAddress &&
		v.StakerNodeId == w.StakerNodeId &&
		v.StakerNodeInfo == w.StakerNodeInfo
}

// StakerAccountRow allows to build a new StakerAccountRow
func NewStakerAccountRow(
	accountAddress string,
	stakerNodeId string,
	stakerNodeInfo string) StakerAccountRow {
	return StakerAccountRow{
		AccountAddress: accountAddress,
		StakerNodeId:   stakerNodeId,
		StakerNodeInfo: stakerNodeInfo,
	}
}

// AccountKeyListRow represents a single row of the account_key_list table
type AccountKeyListRow struct {
	Address        string `db:"address"`
	Index          int    `db:"index"`
	Weight         int    `db:"weight"`
	Revoked        bool   `db:"revoked"`
	SigAlgo        string `db:"sig_algo"`
	HashAlgo       string `db:"hash_algo"`
	PublicKey      string `db:"public_key"`
	SequenceNumber uint64 `db:"sequence_number"`
}

// Equal tells whether v and w represent the same rows
func (v AccountKeyListRow) Equal(w AccountKeyListRow) bool {
	return v.Address == w.Address &&
		v.Index == w.Index &&
		v.Weight == w.Weight &&
		v.Revoked == w.Revoked &&
		v.SigAlgo == w.SigAlgo &&
		v.HashAlgo == w.HashAlgo &&
		v.PublicKey == w.PublicKey &&
		v.SequenceNumber == w.SequenceNumber
}

// AccountKeyListRow allows to build a new AccountKeyListRow
func NewAccountKeyListRow(
	address string,
	index int,
	weight int,
	revoked bool,
	sigAlgo string,
	hashAlgo string,
	publicKey string,
	sequenceNumber uint64) AccountKeyListRow {
	return AccountKeyListRow{
		Address:        address,
		Index:          index,
		Weight:         weight,
		Revoked:        revoked,
		SigAlgo:        sigAlgo,
		HashAlgo:       hashAlgo,
		PublicKey:      publicKey,
		SequenceNumber: sequenceNumber,
	}
}

// StakerNodeIdRow represents a single row of the staker_node_id table
type StakerNodeIdRow struct {
	Address string `db:"address"`
	NodeId  string `db:"node_id"`
}

// Equal tells whether v and w represent the same rows
func (v StakerNodeIdRow) Equal(w StakerNodeIdRow) bool {
	return v.Address == w.Address &&
		v.NodeId == w.NodeId
}

// StakerNodeIdRow allows to build a new StakerNodeIdRow
func NewStakerNodeIdRow(
	address string,
	nodeId string) StakerNodeIdRow {
	return StakerNodeIdRow{
		Address: address,
		NodeId:  nodeId,
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
