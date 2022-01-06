package consensus

import (
	"fmt"

	"github.com/forbole/egldjuno/types"
	"github.com/onflow/flow-go-sdk"

	"github.com/forbole/egldjuno/client"
	db "github.com/forbole/egldjuno/db/postgresql"
	"github.com/rs/zerolog/log"
)

func HandleGenesis(block *flow.Block, chainID string, db *db.Db, flowClient client.Proxy) error {
	log.Debug().Str("module", "consensus").Msg("parsing genesis")
	// Save the genesis time
	err := db.SaveGenesis(types.NewGenesis(block.Timestamp, int64(block.Height), chainID))

	if err != nil {
		return fmt.Errorf("error while storing genesis time: %s", err)
	}

	return nil
}
