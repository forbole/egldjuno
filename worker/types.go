package worker

import (
	"github.com/cosmos/cosmos-sdk/simapp/params"

	"github.com/forbole/egldjuno/client"
	"github.com/forbole/egldjuno/db"
	"github.com/forbole/egldjuno/logging"

	"github.com/forbole/egldjuno/modules/modules"
	"github.com/forbole/egldjuno/types"
)

type Config struct {
	EncodingConfig *params.EncodingConfig
	Queue          types.HeightQueue
	ClientProxy    *client.Proxy
	Database       db.Database
	Modules        []modules.Module
	Logger         logging.Logger
}

func NewConfig(
	queue types.HeightQueue,
	encodingConfig *params.EncodingConfig,
	clientProxy *client.Proxy,
	db db.Database,
	modules []modules.Module,
	logger logging.Logger,
) *Config {
	return &Config{
		EncodingConfig: encodingConfig,
		Queue:          queue,
		ClientProxy:    clientProxy,
		Database:       db,
		Modules:        modules,
		Logger:         logger,
	}
}
