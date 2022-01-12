package auth

import (
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/forbole/egldjuno/modules/messages"
	"github.com/forbole/egldjuno/modules/modules"

	"github.com/forbole/egldjuno/client"
	db "github.com/forbole/egldjuno/db/postgresql"
	"github.com/go-co-op/gocron"

)

var (
	_ modules.Module            = &Module{}
	_ modules.AdditionalOperationsModule = &Module{}


)

// Module represents the x/auth module
type Module struct {
	messagesParser messages.MessageAddressesParser
	encodingConfig *params.EncodingConfig
	clients     client.Proxy
	db             *db.Db
}

// NewModule builds a new Module instance
func NewModule(
	messagesParser messages.MessageAddressesParser,
	clients client.Proxy,
	encodingConfig *params.EncodingConfig, db *db.Db,
) *Module {
	return &Module{
		messagesParser: messagesParser,
		encodingConfig: encodingConfig,
		clients:     clients,
		db:             db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "auth"
}

// RegisterPeriodicOperations implements modules.Module
func (m *Module) RunAdditionalOperations() error {
	return RunAdditionalOperations( m.db, m.clients)
}

// RegisterPeriodicOperations implements modules.Module
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	return RegisterPeriodicOperations(scheduler, m.db,m.clients)
}