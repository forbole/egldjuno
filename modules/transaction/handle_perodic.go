package transaction

import (
	"github.com/forbole/egldjuno/client"
	db "github.com/forbole/egldjuno/db/postgresql"
	txutils "github.com/forbole/egldjuno/modules/transaction/utils"
	"github.com/forbole/egldjuno/modules/utils"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

// Register registers the utils that should be run periodically
func RegisterPeriodicOperations(scheduler *gocron.Scheduler, db *db.Db, client client.Proxy) error {
	log.Debug().Str("module", "tx").Msg("setting up periodic tasks")

	if _, err := scheduler.Every(1).Second().Do(func() {
		utils.WatchMethod(func() error { return getNewTransactions(db, client) })
	}); err != nil {
		return err
	}

	return nil
}

func getNewTransactions(db *db.Db, client client.Proxy) error {
	txs, err := txutils.GetNewTransactions(client)
	if err != nil {
		return err
	}
	return db.SaveTxs(txs)
}
