package utils

import (
	"os"

	"github.com/rs/zerolog"
)

func GetLogger() zerolog.Logger {
	return zerolog.
		New(os.Stderr).
		Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
