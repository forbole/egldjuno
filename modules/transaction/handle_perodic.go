package transaction

import (
	"fmt"

	"github.com/forbole/egldjuno/client"
	db "github.com/forbole/egldjuno/db/postgresql"
	txutils "github.com/forbole/egldjuno/modules/transaction/utils"
	"github.com/forbole/egldjuno/modules/utils"
	"github.com/forbole/egldjuno/types"
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

	for _, tx := range txs {
		var txResults []types.SmartContractResult
		for _, smr := range tx.SmartContractResult {
			fmt.Println(smr.Logs)
			fmt.Println(smr.Hash)

			txResults = append(txResults, smr)
		}
		err = db.SaveSmartContractResult(txResults, tx.TxHash)
		if err != nil {
			return fmt.Errorf("Error when saving smart_contract)result into db:%s", err)
		}
	}

	err = db.SaveTxs(txs)
	if err != nil {
		return fmt.Errorf("Error when saving tx:%s", err)
	}
	return nil
}
