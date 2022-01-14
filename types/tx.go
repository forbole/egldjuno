package types

import "reflect"

// Tx represents an already existing blockchain transaction
type Txs []Tx
type Tx struct {
	TxHash          string
	GasLimit        int64
	GasPrice        int64
	GasUsed         int64
	MiniBlockHash   string
	Nonce           int64
	Receiver        string
	ReceiverShard   int64
	Round           int64
	Sender          string
	SenderShard     int64
	Signature       string
	Status          string
	Value           string
	Fee             string
	Timestamp       int64
	Data            string
	SmartContractResult []SmartContractResult `json:"results"`
}

// Equal tells whether v and w represent the same rows
func (v Tx) Equal(w Tx) bool {
	return v.TxHash == w.TxHash &&
		v.GasLimit == w.GasLimit &&
		v.GasPrice == w.GasPrice &&
		v.GasUsed == w.GasUsed &&
		v.MiniBlockHash == w.MiniBlockHash &&
		v.Nonce == w.Nonce &&
		v.Receiver == w.Receiver &&
		v.ReceiverShard == w.ReceiverShard &&
		v.Round == w.Round &&
		v.Sender == w.Sender &&
		v.SenderShard == w.SenderShard &&
		v.Signature == w.Signature &&
		v.Status == w.Status &&
		v.Value == w.Value &&
		v.Fee == w.Fee &&
		v.Timestamp == w.Timestamp &&
		v.Data == w.Data 
}

// Transaction allows to build a new Transaction
func NewTx(
	txHash string,
	gasLimit int64,
	gasPrice int64,
	gasUsed int64,
	miniBlockHash string,
	nonce int64,
	receiver string,
	receiverShard int64,
	round int64,
	sender string,
	senderShard int64,
	signature string,
	status string,
	value string,
	fee string,
	timestamp int64,
	data string) Tx {
	return Tx{
		TxHash:          txHash,
		GasLimit:        gasLimit,
		GasPrice:        gasPrice,
		GasUsed:         gasUsed,
		MiniBlockHash:   miniBlockHash,
		Nonce:           nonce,
		Receiver:        receiver,
		ReceiverShard:   receiverShard,
		Round:           round,
		Sender:          sender,
		SenderShard:     senderShard,
		Signature:       signature,
		Status:          status,
		Value:           value,
		Fee:             fee,
		Timestamp:       timestamp,
		Data:            data,
	}
}

type Action struct {
	Category    string `json:"category"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Arguments   []byte `json:"arguments"`
}

type SmartContractResult struct { 
		Hash           string `json:"hash"`
		Timestamp      int    `json:"timestamp"`
		Nonce          int    `json:"nonce"`
		GasLimit       int    `json:"gasLimit"`
		GasPrice       int    `json:"gasPrice"`
		Value          string `json:"value"`
		Sender         string `json:"sender"`
		Receiver       string `json:"receiver"`
		Data           string `json:"data"`
		PrevTxHash     string `json:"prevTxHash"`
		OriginalTxHash string `json:"originalTxHash"`
		CallType       string `json:"callType"`
		RelayedValue *string `json:"relayedValue"`
	Logs []byte `json:"logs"`
	ReturnMessage []byte `json:"returnMessage"`
  }
  
  // Equal tells whether v and w represent the same rows
  func (v SmartContractResult) Equal(w SmartContractResult)bool{
	return v.Hash==w.Hash && 
  v.Timestamp==w.Timestamp && 
  v.Nonce==w.Nonce && 
  v.GasLimit==w.GasLimit && 
  v.GasPrice==w.GasPrice && 
  v.Value==w.Value && 
  v.Sender==w.Sender && 
  v.Receiver==w.Receiver && 
  v.RelayedValue==w.RelayedValue && 
  v.Data==w.Data && 
  v.PrevTxHash==w.PrevTxHash && 
  v.OriginalTxHash==w.OriginalTxHash && 
  v.CallType==w.CallType && 
  reflect.DeepEqual(v.Logs,w.Logs) &&
reflect.DeepEqual(v.ReturnMessage,w.ReturnMessage) }
  
   // SmartContractResult allows to build a new SmartContractResult
  func NewSmartContractResult( 
	txHash string,
	hash string,
	timestamp int,
	nonce int,
	gasLimit int,
	gasPrice int,
	value string,
	sender string,
	receiver string,
	relayedValue string,
	data string,
	prevTxHash string,
	originalTxHash string,
	callType string,
	logs []byte,
	returnMessage []byte) SmartContractResult{
   return SmartContractResult{
   Hash:hash,
   Timestamp:timestamp,
   Nonce:nonce,
   GasLimit:gasLimit,
   GasPrice:gasPrice,
   Value:value,
   Sender:sender,
   Receiver:receiver,
   RelayedValue:&relayedValue,
   Data:data,
   PrevTxHash:prevTxHash,
   OriginalTxHash:originalTxHash,
   CallType:callType,
   Logs:logs,
   ReturnMessage: returnMessage,
  }
  }