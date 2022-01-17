package types

// AccountRow represents a single row of the account table
// AccountRow represents a single row of the account table
type AccountRow struct {
	Address         string `db:"address"`
	Nonce           int64  `db:"nonce"`
	Balance         string `db:"balance"`
	RootHash        string `db:"root_hash"`
	TxCount         int64  `db:"tx_count"`
	ScrCount        int64  `db:"scr_count"`
	Shard           int64  `db:"shard"`
	DeveloperReward string `db:"developer_reward"`
}

// Equal tells whether v and w represent the same rows
func (v AccountRow) Equal(w AccountRow) bool {
	return v.Address == w.Address &&
		v.Nonce == w.Nonce &&
		v.Balance == w.Balance &&
		v.RootHash == w.RootHash &&
		v.TxCount == w.TxCount &&
		v.ScrCount == w.ScrCount &&
		v.Shard == w.Shard &&
		v.DeveloperReward == w.DeveloperReward
}

// AccountRow allows to build a new AccountRow
func NewAccountRow(
	address string,
	nonce int64,
	balance string,
	rootHash string,
	txCount int64,
	scrCount int64,
	shard int64,
	developerReward string) AccountRow {
	return AccountRow{
		Address:         address,
		Nonce:           nonce,
		Balance:         balance,
		RootHash:        rootHash,
		TxCount:         txCount,
		ScrCount:        scrCount,
		Shard:           shard,
		DeveloperReward: developerReward,
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
