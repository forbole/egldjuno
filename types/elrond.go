package types

import (
	"time"
)

// Validator contains the data of a single validator
type Validator struct {
	ConsAddr   string
	ConsPubKey string
}

// NewValidator allows to build a new Validator instance
func NewValidator(consAddr string, consPubKey string) *Validator {
	return &Validator{
		ConsAddr:   consAddr,
		ConsPubKey: consPubKey,
	}
}

// -------------------------------------------------------------------------------------------------------------------

// CommitSig contains the data of a single validator commit signature
type CommitSig struct {
	Height           int64
	ValidatorAddress string
	VotingPower      int64
	ProposerPriority int64
	Timestamp        time.Time
}

// NewCommitSig allows to build a new CommitSign object
func NewCommitSig(validatorAddress string, votingPower, proposerPriority, height int64, timestamp time.Time) *CommitSig {
	return &CommitSig{
		Height:           height,
		ValidatorAddress: validatorAddress,
		VotingPower:      votingPower,
		ProposerPriority: proposerPriority,
		Timestamp:        timestamp,
	}
}

// -------------------------------------------------------------------------------------------------------------------

// Block contains the data of a single chain block
type Block struct { 
	Hash string
	Epoch int64
	Nonce int64
	PrevHash string
	Proposer int64
	PubKeyBitmap string
	Round int64
	Shard int64
	Size int64
	SizeTxs int64
	StateRootHash string
	TimeStamp int64
	TxCount int64
	GasConsumed int64
	GasRefunded int64
	GasPenalized int64
	MaxGasLimit int64
	
	MiniBlocksHashes []string
	Validators []int64
	NotarizedBlocksHashes []string
  }
// Equal tells whether v and w represent the same rows
func (v Block) Equal(w Block)bool{
	return v.Hash==w.Hash && 
  v.Epoch==w.Epoch && 
  v.Nonce==w.Nonce && 
  v.PrevHash==w.PrevHash && 
  v.Proposer==w.Proposer && 
  v.PubKeyBitmap==w.PubKeyBitmap && 
  v.Round==w.Round && 
  v.Shard==w.Shard && 
  v.Size==w.Size && 
  v.SizeTxs==w.SizeTxs && 
  v.StateRootHash==w.StateRootHash && 
  v.TimeStamp==w.TimeStamp && 
  v.TxCount==w.TxCount && 
  v.GasConsumed==w.GasConsumed && 
  v.GasRefunded==w.GasRefunded && 
  v.GasPenalized==w.GasPenalized && 
  v.MaxGasLimit==w.MaxGasLimit }
  
   // Block allows to build a new Block
  func NewBlock( 
	hash string,
	epoch int64,
	nonce int64,
	prevHash string,
	proposer int64,
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
	maxGasLimit int64) Block{
   return Block{
   Hash:hash,
   Epoch:epoch,
   Nonce:nonce,
   PrevHash:prevHash,
   Proposer:proposer,
   PubKeyBitmap:pubKeyBitmap,
   Round:round,
   Shard:shard,
   Size:size,
   SizeTxs:sizeTxs,
   StateRootHash:stateRootHash,
   TimeStamp:timeStamp,
   TxCount:txCount,
   GasConsumed:gasConsumed,
   GasRefunded:gasRefunded,
   GasPenalized:gasPenalized,
   MaxGasLimit:maxGasLimit,
  }
}

// Message represents the data of a single message
type Message struct {
	TxHash    string
	Index     int
	Type      string
	Value     string
	Addresses []string
}

// NewMessage allows to build a new Message instance
func NewMessage(txHash string, index int, msgType string, value string, addresses []string) *Message {
	return &Message{
		TxHash:    txHash,
		Index:     index,
		Type:      msgType,
		Value:     value,
		Addresses: addresses,
	}
}
