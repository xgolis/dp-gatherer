package main

import (
	"os"

	"github.com/xgolis/dp-gatherer/cmd/gatherer/app"
	"github.com/xgolis/dp-gatherer/cmd/gatherer/pkg/errhelp"
	"github.com/xgolis/dp-gatherer/cmd/gatherer/pkg/utils"
)

func main() {
	err := app.New().Execute()
	// logger := zerolog.
	//  New(os.Stderr).
	//  Output(zerolog.ConsoleWriter{Out: os.Stderr})
	logger := utils.GetLogger()
	if err != nil {
		logger.Err(err).Msg("Program exited with error")
		if errDescription := errhelp.GetHelp(err); errDescription != "" {
			logger.Error().Msg(errDescription)
		}
		os.Exit(1)
	}

}
