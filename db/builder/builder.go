package builder

import (
	"github.com/cosmos/cosmos-sdk/simapp/params"

	"github.com/HarleyAppleChoi/junomum/types"

	"github.com/HarleyAppleChoi/junomum/db"

	database "github.com/HarleyAppleChoi/junomum/db/postgresql"
)

// Builder represents a generic Builder implementation that build the proper database
// instance based on the configuration the user has specified
func Builder(cfg types.Config, encodingConfig *params.EncodingConfig) (db.Database, error) {
	return database.Builder(cfg, encodingConfig)
}
