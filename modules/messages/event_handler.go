package messages

/*
import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"

	"github.com/forbole/egldjuno/db"
	"github.com/forbole/egldjuno/types"
)

// HandleEvent represents a message handler that stores the given message inside the proper database table
func HandleEvent(
	index int, msg sdk.Msg, tx *types.Txs,
	parseAddresses MessageAddressesParser, cdc codec.Marshaler, db db.Database,
) error {
	// Get the involved addresses
	addresses, err := parseAddresses(cdc, msg)
	if err != nil {
		return err
	}

	// Marshal the value properly
	bz, err := cdc.MarshalJSON(msg)
	if err != nil {
		return err
	}

	return db.SaveMessage(types.NewMessage(
		tx.TxHash,
		index,
		proto.MessageName(msg),
		string(bz),
		addresses,
	))
}
*/
