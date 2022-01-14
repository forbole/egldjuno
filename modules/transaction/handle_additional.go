package transaction

import (
	"github.com/forbole/egldjuno/client"
	db "github.com/forbole/egldjuno/db/postgresql"
	//txutils "github.com/forbole/egldjuno/modules/transaction/utils"
	//"github.com/forbole/egldjuno/modules/utils"
)

func RunAdditionalOperations(db *db.Db, client client.Proxy) error {
	//utils.WatchMethod(func() error { return txutils.GetHistoricBlocks(db, client) })
	return nil
}
