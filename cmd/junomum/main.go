package main

import (
	"os"

	"github.com/HarleyAppleChoi/junomum/cmd/parse"

	"github.com/HarleyAppleChoi/junomum/modules/messages"

	"github.com/HarleyAppleChoi/junomum/cmd"
	database "github.com/HarleyAppleChoi/junomum/db/postgresql"
	"github.com/HarleyAppleChoi/junomum/modules"
	"github.com/HarleyAppleChoi/junomum/types/config"
)

func main() {
	// Config the runner
	config := cmd.NewConfig("junomum").
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
