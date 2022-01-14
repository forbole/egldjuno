package types

// SmartContractResultRow represents a single row of the smart_contract_result table
type SmartContractResultRow struct {
	TxHash         string `db:"tx_hash"`
	Hash           string `db:"hash"`
	Timestamp      int64  `db:"timestamp"`
	Nonce          int64  `db:"nonce"`
	GasLimit       int64  `db:"gas_limit"`
	GasPrice       int64  `db:"gas_price"`
	Value          string `db:"value"`
	Sender         string `db:"sender"`
	Receiver       string `db:"receiver"`
	RelayedValue   string `db:"relayed_value"`
	Data           string `db:"data"`
	PrevTxHash     string `db:"prev_tx_hash"`
	OriginalTxHash string `db:"original_tx_hash"`
	CallType       string `db:"call_type"`
	Logs           string `db:"logs"`
}

// Equal tells whether v and w represent the same rows
func (v SmartContractResultRow) Equal(w SmartContractResultRow) bool {
	return v.TxHash == w.TxHash &&
		v.Hash == w.Hash &&
		v.Timestamp == w.Timestamp &&
		v.Nonce == w.Nonce &&
		v.GasLimit == w.GasLimit &&
		v.GasPrice == w.GasPrice &&
		v.Value == w.Value &&
		v.Sender == w.Sender &&
		v.Receiver == w.Receiver &&
		v.RelayedValue == w.RelayedValue &&
		v.Data == w.Data &&
		v.PrevTxHash == w.PrevTxHash &&
		v.OriginalTxHash == w.OriginalTxHash &&
		v.CallType == w.CallType &&
		v.Logs == w.Logs
}

// SmartContractResultRow allows to build a new SmartContractResultRow
func NewSmartContractResultRow(
	txHash string,
	hash string,
	timestamp int64,
	nonce int64,
	gasLimit int64,
	gasPrice int64,
	value string,
	sender string,
	receiver string,
	relayedValue string,
	data string,
	prevTxHash string,
	originalTxHash string,
	callType string,
	logs string) SmartContractResultRow {
	return SmartContractResultRow{
		TxHash:         txHash,
		Hash:           hash,
		Timestamp:      timestamp,
		Nonce:          nonce,
		GasLimit:       gasLimit,
		GasPrice:       gasPrice,
		Value:          value,
		Sender:         sender,
		Receiver:       receiver,
		RelayedValue:   relayedValue,
		Data:           data,
		PrevTxHash:     prevTxHash,
		OriginalTxHash: originalTxHash,
		CallType:       callType,
		Logs:           logs,
	}
}
