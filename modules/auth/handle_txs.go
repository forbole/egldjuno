package auth

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/egldjuno/modules/messages"
	"github.com/forbole/egldjuno/types"

	"github.com/forbole/egldjuno/client"
	db "github.com/forbole/egldjuno/db/postgresql"
	authutils "github.com/forbole/egldjuno/modules/auth/utils"
)

// HandleEvent handles any message updating the involved accounts
func HandleTxs(getAddresses messages.MessageAddressesParser, cdc codec.Marshaler, db *db.Db, flowClient client.Proxy, tx *types.Tx) error {
	height, err := flowClient.LatestHeight()
	if err != nil {
		return err
	}

	addresses, err := getAddresses(cdc, *tx)
	if err != nil {
		return err
	}

	return authutils.UpdateAccounts(addresses, db, height, flowClient)

}
