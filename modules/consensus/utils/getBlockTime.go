package utils

import (
	database "github.com/forbole/egldjuno/db/postgresql"
	"github.com/forbole/egldjuno/types"
	"github.com/onflow/flow-go-sdk"
)

// GetBlockTimeInMinute get last block in db and get the block time in minute
func GetBlockTimeInMinute(db *database.Db) (*types.BlockTime, error) {
	/*
		block, err := db.GetLastBlock()
		if err != nil {
			return nil, err
		}

		genesis, err := db.GetGenesis()
		if err != nil {
			return nil, err
		} */

	/* // Check if the chain has been created at least a minute ago
	if block.TimeStamp.Sub(genesis.Time).Minutes() < 0 {
		return nil, nil
	}

	minute, err := db.GetBlockHeightTimeMinuteAgo(block.Timestamp)
	if err != nil {
		return nil, err
	}
	newBlockTime := block.Timestamp.Sub(minute.Timestamp).Seconds() / float64(block.Height-minute.Height)

	blocktime := types.NewBlockTime(block.Height, newBlockTime)
	return &blocktime, nil */

	return nil, nil
}

// GetBlockTimeInHour get last block in db and get the block time in hour

func GetBlockTimeInHour(db *database.Db) (*types.BlockTime, error) {

	/* block, err := db.GetLastBlock()
	if err != nil {
		return nil, fmt.Errorf("fail to get block time in hour:%s", err)
	}

	genesis, err := db.GetGenesis()
	if err != nil {
		return nil, fmt.Errorf("fail to get block time in hour:%s", err)
	}

	// Check if the chain has been created at least an hour ago
	if block.Timestamp.Sub(genesis.Time).Hours() < 0 {
		return nil, fmt.Errorf("fail to get block time in hour:%s", err)
	}

	hour, err := db.GetBlockHeightTimeHourAgo(block.Timestamp)
	if err != nil {
		return nil, fmt.Errorf("fail to get block time in hour:%s", err)
	}
	newBlockTime := block.Timestamp.Sub(hour.Timestamp).Seconds() / float64(block.Height-hour.Height)

	blocktime := types.NewBlockTime(block.Height, newBlockTime)

	return &blocktime, nil */
	return nil, nil

}

// GetBlockTimeInDay get last block in db and get the block time in day

func GetBlockTimeInDay(db *database.Db) (*types.BlockTime, error) {

	/* block, err := db.GetLastBlock()
	if err != nil {
		return nil, fmt.Errorf("fail to get block time in days:%s", err)
	}

	genesis, err := db.GetGenesis()
	if err != nil {
		return nil, fmt.Errorf("fail to get block time in days:%s", err)
	}

	// Check if the chain has been created at least a days ago
	if block.Timestamp.Sub(genesis.Time).Hours() < 24 {
		return nil, fmt.Errorf("fail to get block time in days:%s", err)
	}

	day, err := db.GetBlockHeightTimeDayAgo(block.Timestamp)
	if err != nil {
		return nil, fmt.Errorf("fail to get block time in days:%s", err)
	}
	newBlockTime := block.Timestamp.Sub(day.Timestamp).Seconds() / float64(block.Height-day.Height)
	blocktime := types.NewBlockTime(block.Height, newBlockTime)

	return &blocktime, nil */
	return nil, nil
}

// GetGenesisBlockTime get the genesis block from db and calculate block time using latest block
func GetGenesisBlockTime(db *database.Db, block flow.Block) (*types.BlockTime, error) {
	genesis, err := db.GetGenesis()
	if err != nil {
		return nil, err
	}

	newBlockTime := block.Timestamp.Sub(genesis.Time).Seconds() / float64(int64(block.Height)-genesis.InitialHeight)

	blocktime := types.NewBlockTime(int64(block.Height), newBlockTime)
	return &blocktime, nil

}
