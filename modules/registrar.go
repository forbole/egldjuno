package modules

import (
	"github.com/HarleyAppleChoi/junomum/client"
	"github.com/HarleyAppleChoi/junomum/db"
	"github.com/HarleyAppleChoi/junomum/db/postgresql"
	"github.com/cosmos/cosmos-sdk/simapp/params"

	"github.com/HarleyAppleChoi/junomum/modules/messages"
	"github.com/HarleyAppleChoi/junomum/modules/modules"
	"github.com/HarleyAppleChoi/junomum/modules/registrar"
	"github.com/HarleyAppleChoi/junomum/types"

	"github.com/HarleyAppleChoi/junomum/modules/auth"
	"github.com/HarleyAppleChoi/junomum/modules/consensus"
	"github.com/HarleyAppleChoi/junomum/modules/telemetry"
)

var (
	_ registrar.Registrar = &Registrar{}
)

// Registrar represents the modules.Registrar that allows to register all modules that are supported by BigDipper
type Registrar struct {
	parser messages.MessageAddressesParser
}

// NewRegistrar allows to build a new Registrar instance
func NewRegistrar(parser messages.MessageAddressesParser) *Registrar {
	return &Registrar{
		parser: parser,
	}
}

// BuildModules implements modules.Registrar
func (r *Registrar) BuildModules(
	cfg types.Config, encodingConfig *params.EncodingConfig, database db.Database, cp *client.Proxy,
) modules.Modules {

	bigDipperBd := postgresql.Cast(database)

	return []modules.Module{
		messages.NewModule(r.parser, encodingConfig.Marshaler, database),
		auth.NewModule(r.parser, *cp, encodingConfig, bigDipperBd),
		consensus.NewModule(r.parser, *cp, encodingConfig, bigDipperBd),
		telemetry.NewModule(cfg, r.parser, *cp, encodingConfig, bigDipperBd),
	}
}
