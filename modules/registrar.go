package modules

import (
	"github.com/forbole/egldjuno/client"
	"github.com/forbole/egldjuno/db"
	"github.com/forbole/egldjuno/db/postgresql"
	"github.com/cosmos/cosmos-sdk/simapp/params"

	"github.com/forbole/egldjuno/modules/messages"
	"github.com/forbole/egldjuno/modules/modules"
	"github.com/forbole/egldjuno/modules/registrar"
	"github.com/forbole/egldjuno/types"

	"github.com/forbole/egldjuno/modules/auth"
	"github.com/forbole/egldjuno/modules/consensus"
	"github.com/forbole/egldjuno/modules/telemetry"
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
