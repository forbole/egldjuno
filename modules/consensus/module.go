package consensus

import (
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/HarleyAppleChoi/junomum/modules/messages"
	"github.com/HarleyAppleChoi/junomum/modules/modules"
	"github.com/HarleyAppleChoi/junomum/types"
	"github.com/go-co-op/gocron"
	"github.com/onflow/flow-go-sdk"

	"github.com/HarleyAppleChoi/junomum/client"
	db "github.com/HarleyAppleChoi/junomum/db/postgresql"
)

var (
	_ modules.Module      = &Module{}
	_ modules.BlockModule = &Module{}
)

// Module represents the x/auth module
type Module struct {
	messagesParser messages.MessageAddressesParser
	encodingConfig *params.EncodingConfig
	flowClient     client.Proxy
	db             *db.Db
}

// NewModule builds a new Module instance
func NewModule(
	messagesParser messages.MessageAddressesParser,
	flowClient client.Proxy,
	encodingConfig *params.EncodingConfig, db *db.Db,
) *Module {
	return &Module{
		messagesParser: messagesParser,
		encodingConfig: encodingConfig,
		flowClient:     flowClient,
		db:             db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "consensus"
}

// HandleEvent implements modules.MessageModule
func (m *Module) HandleBlock(block *flow.Block, _ *types.Txs) error {
	return HandleBlock(block, m.messagesParser, m.db, int64(block.Height), m.flowClient)
}

// RegisterPeriodicOperations implements modules.Module
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	return Register(scheduler, m.db)
}

// HandleGenesis implements modules.Module
func (m *Module) HandleGenesis(block *flow.Block, chainID string) error {
	return HandleGenesis(block, chainID, m.db, m.flowClient)
}
