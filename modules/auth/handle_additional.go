package auth

import (
	db "github.com/forbole/egldjuno/db/postgresql"
	"github.com/forbole/egldjuno/client"
	"github.com/forbole/egldjuno/modules/utils"
	authutils "github.com/forbole/egldjuno/modules/auth/utils"
)

func RunAdditionalOperations( db *db.Db, client client.Proxy)error{
	utils.WatchMethod(func() error { return authutils.GetHistoricBlocks(db,client) })
	return nil
}
