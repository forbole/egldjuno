package auth

import (
	"github.com/forbole/egldjuno/client"
	db "github.com/forbole/egldjuno/db/postgresql"
	authutils "github.com/forbole/egldjuno/modules/auth/utils"
	"github.com/forbole/egldjuno/modules/utils"
)

func RunAdditionalOperations(db *db.Db, client client.Proxy) error {
	utils.WatchMethod(func() error { return authutils.GetHistoricBlocks(db, client) })
	return nil
}
