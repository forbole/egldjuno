package telemetry

import (
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/HarleyAppleChoi/junomum/modules/messages"
	"github.com/HarleyAppleChoi/junomum/modules/modules"

	"github.com/HarleyAppleChoi/junomum/client"
	"github.com/HarleyAppleChoi/junomum/types"

	db "github.com/HarleyAppleChoi/junomum/db/postgresql"
)

const (
	ModuleName = "telemetry"
)

var (
	_ modules.Module                     = &Module{}
	_ modules.AdditionalOperationsModule = &Module{}
)

// Module represents the telemetry module
type Module struct {
	cfg            types.Config
	messagesParser messages.MessageAddressesParser
	encodingConfig *params.EncodingConfig
	flowClient     client.Proxy
	db             *db.Db
}

// NewModule returns a new Module implementation
func NewModule(
	cfg types.Config,
	messagesParser messages.MessageAddressesParser,
	flowClient client.Proxy,
	encodingConfig *params.EncodingConfig, db *db.Db,
) *Module {
	return &Module{
		cfg:            cfg,
		messagesParser: messagesParser,
		encodingConfig: encodingConfig,
		flowClient:     flowClient,
		db:             db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return ModuleName
}

// RunAdditionalOperations implements modules.AdditionalOperationsModule
func (m *Module) RunAdditionalOperations() error {
	return RunAdditionalOperations(m.cfg)
}
