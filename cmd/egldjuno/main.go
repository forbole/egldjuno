package main

import (
	"os"

	"github.com/forbole/egldjuno/cmd/parse"

	"github.com/forbole/egldjuno/modules/messages"

	"github.com/forbole/egldjuno/cmd"
	database "github.com/forbole/egldjuno/db/postgresql"
	"github.com/forbole/egldjuno/modules"
	"github.com/forbole/egldjuno/types/config"
)

func main() {
	// Config the runner
	config := cmd.NewConfig("egldjuno").
		WithParseConfig(parse.NewConfig().
			WithConfigParser(config.ParseConfig).
			WithDBBuilder(database.Builder).
			WithEncodingConfigBuilder(config.MakeEncodingConfig(nil)).
			WithRegistrar(modules.NewRegistrar(getAddressesParser())),
		)

	// Run the commands and panic on any error
	exec := cmd.BuildDefaultExecutor(config)
	err := exec.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func getAddressesParser() messages.MessageAddressesParser {
	return messages.DefaultMessagesParser
}
