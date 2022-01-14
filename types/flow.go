package types

import (
	"reflect"

	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
)

type Event struct {
	//Transaction Result Event
	Height           int
	Type             string
	TransactionID    string
	TransactionIndex int
	EventIndex       int
	Value            cadence.Event
}

func NewEvent(height int, t string, transactionID string, transactionIndex int, eventIndex int,
	value cadence.Event) Event {

	return Event{
		//Transaction Result Event
		Height:           height,
		Type:             t,
		TransactionID:    transactionID,
		TransactionIndex: transactionIndex,
		EventIndex:       eventIndex,
		Value:            value,
	}
}

// Successful tells whether this tx is successful or not
func (tx Tx) Successful() bool {
	return true
}

type Collection struct {
	Height         uint64
	Id             string
	Processed      bool
	TransactionIds []flow.Identifier
}

// Equal tells whether v and w represent the same rows
func (v Collection) Equal(w Collection) bool {
	return v.Height == w.Height &&
		v.Id == w.Id &&
		v.Processed == w.Processed &&
		reflect.DeepEqual(v.TransactionIds, w.TransactionIds)
}

// Collection allows to build a new Collection
func NewCollection(
	height uint64,
	id string,
	processed bool,
	transactionIds []flow.Identifier) Collection {

	return Collection{
		Height:         height,
		Id:             id,
		Processed:      processed,
		TransactionIds: transactionIds,
	}
}

type TransactionResult struct {
	TransactionId string
	Status        string
	Error         string
}

// Equal tells whether v and w represent the same rows
func (v TransactionResult) Equal(w TransactionResult) bool {
	return v.TransactionId == w.TransactionId &&
		v.Status == w.Status &&
		v.Error == w.Error
}

// TransactionResult allows to build a new TransactionResult
func NewTransactionResult(
	transactionId string,
	status string,
	error string) TransactionResult {
	return TransactionResult{
		TransactionId: transactionId,
		Status:        status,
		Error:         error,
	}
}
