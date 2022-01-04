package logging

import (
	"github.com/onflow/flow-go-sdk"

	"github.com/HarleyAppleChoi/junomum/modules/modules"
	"github.com/HarleyAppleChoi/junomum/types"
)

const (
	LogKeyModule  = "module"
	LogKeyHeight  = "height"
	LogKeyTxHash  = "tx_hash"
	LogKeyMsgType = "msg_type"
)

// Logger defines a function that takes an error and logs it.
type Logger interface {
	SetLogLevel(level string) error
	SetLogFormat(format string) error

	Info(msg string, keyvals ...interface{})
	Debug(msg string, keyvals ...interface{})
	Error(msg string, keyvals ...interface{})

	GenesisError(module modules.Module, err error)
	BlockError(module modules.Module, block *flow.Block, err error)
	EventsError(module modules.Module, results *types.Event, err error)
	TxError(module modules.Module, tx *types.Tx, err error)
}
