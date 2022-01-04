package auth

import (
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"

	"github.com/HarleyAppleChoi/junomum/client"
	"github.com/HarleyAppleChoi/junomum/modules/utils"

	database "github.com/HarleyAppleChoi/junomum/db/postgresql"
	db "github.com/HarleyAppleChoi/junomum/db/postgresql"
	authutils "github.com/HarleyAppleChoi/junomum/modules/auth/utils"
)

func RegisterPeriodicOps(scheduler *gocron.Scheduler, db *database.Db, flowClient client.Proxy) error {
	log.Debug().Str("module", "staking").Msg("setting up periodic tasks")

	if _, err := scheduler.Every(1).Week().Tuesday().At("15:00").StartImmediately().Do(func() {
		utils.WatchMethod(func() error { return HandleAccounts(db, flowClient) })
	}); err != nil {
		return err
	}

	return HandleAccounts(db, flowClient)
}

func HandleAccounts(db *db.Db, flowClient client.Proxy) error {
	//get Accounts
	accountStringArray, err := db.GetAddresses()
	if err != nil {
		return err
	}
	if len(accountStringArray) == 0 {
		return nil
	}

	height, err := flowClient.LatestHeight()
	if err != nil {
		return err
	}

	lockedAccountBalances, err := authutils.GetLockedAccountBalance(accountStringArray, height, flowClient)
	if err != nil {
		return err
	}

	if len(lockedAccountBalances) == 0 {
		return nil
	}

	err = db.SaveLockedAccountBalance(lockedAccountBalances)
	if err != nil {
		return err
	}

	return nil

}
