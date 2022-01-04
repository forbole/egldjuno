package worker

import (
	"github.com/cosmos/cosmos-sdk/simapp/params"

	"github.com/HarleyAppleChoi/junomum/client"
	"github.com/HarleyAppleChoi/junomum/db"
	"github.com/HarleyAppleChoi/junomum/logging"

	"github.com/HarleyAppleChoi/junomum/modules/modules"
	"github.com/HarleyAppleChoi/junomum/types"
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
