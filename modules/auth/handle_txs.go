package auth

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/HarleyAppleChoi/junomum/modules/messages"
	"github.com/HarleyAppleChoi/junomum/types"

	"github.com/HarleyAppleChoi/junomum/client"
	db "github.com/HarleyAppleChoi/junomum/db/postgresql"
	authutils "github.com/HarleyAppleChoi/junomum/modules/auth/utils"
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
