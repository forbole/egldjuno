package auth

import (
	"github.com/forbole/egldjuno/client"
	db "github.com/forbole/egldjuno/db/postgresql"
	authutils "github.com/forbole/egldjuno/modules/auth/utils"
	"github.com/forbole/egldjuno/modules/utils"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

// Register registers the utils that should be run periodically
func RegisterPeriodicOperations(scheduler *gocron.Scheduler, db *db.Db, client client.Proxy) error {
	log.Debug().Str("module", "consensus").Msg("setting up periodic tasks")

	if _, err := scheduler.Every(1).Second().Do(func() {
		utils.WatchMethod(func() error { return getNewBlocks(db, client) })
	}); err != nil {
		return err
	}
	if _, err := scheduler.Every(1).Second().Do(func() {
		utils.WatchMethod(func() error { return getNewAccounts(db, client) })
	}); err != nil {
		return err
	}

	return nil
}

func getNewBlocks(db *db.Db, client client.Proxy) error {
	blocks, err := authutils.GetNewBlocks(client)
	if err != nil {
		return err
	}
	return db.SaveBlock(blocks)
}

func getNewAccounts(db *db.Db, client client.Proxy) error {
	accounts, err := authutils.GetNewAccounts(client)
	if err != nil {
		return err
	}
	return db.SaveAccount(accounts)
}
