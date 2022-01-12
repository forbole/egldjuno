package types

import (
	"time"
)

type GenesisRow struct {
	OneRowID      bool      `db:"one_row_id"`
	Time          time.Time `db:"time"`
	InitialHeight int64     `db:"initial_height"`
	ChainId       string    `db:"chain_id"`
}

func NewGenesisRow(time time.Time, initialHeight int64, chainId string) GenesisRow {
	return GenesisRow{
		OneRowID:      true,
		Time:          time,
		InitialHeight: initialHeight,
		ChainId:       chainId,
	}
}

func (r GenesisRow) Equal(s GenesisRow) bool {
	return r.Time.Equal(s.Time) &&
		r.InitialHeight == s.InitialHeight &&
		r.ChainId == s.ChainId
}

// -------------------------------------------------------------------------------------------------------------------
// AverageTimeRow is the average block time each minute/hour/day
type AverageTimeRow struct {
	OneRowID    bool    `db:"one_row_id"`
	AverageTime float64 `db:"average_time"`
	Height      int64   `db:"height"`
}

func NewAverageTimeRow(averageTime float64, height int64) AverageTimeRow {
	return AverageTimeRow{
		OneRowID:    true,
		AverageTime: averageTime,
		Height:      height,
	}
}

// Equal return true if two AverageTimeRow are true
func (r AverageTimeRow) Equal(s AverageTimeRow) bool {
	return r.AverageTime == s.AverageTime &&
		r.Height == s.Height
}
