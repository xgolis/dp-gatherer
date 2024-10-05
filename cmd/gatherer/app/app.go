package app

import (
	"fmt"
	"os"
	"time"

	"github.com/xgolis/dp-gatherer/git"
	"github.com/xgolis/dp-gatherer/pkg/version"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var ErrorVersion = fmt.Errorf("operating under outdated version")

// createLogger creates a new zerolog.Logger instance with the log level based
// on the specified verbosity.
//
// The logger output is a human-friendly format sent to standard error.
//
// Verbosity log-levels: 0: Info, 1: Debug, 2+: Trace
func createLogger(verbosity int) zerolog.Logger {
	level := zerolog.InfoLevel
	switch verbosity {
	case 0:
		level = zerolog.InfoLevel
	case 1:
		level = zerolog.DebugLevel
	default:
		level = zerolog.TraceLevel
	}
	zerolog.DurationFieldUnit = time.Second

	return zerolog.New(os.Stderr).
		With().Timestamp().
		Logger().
		Level(level).
		Output(zerolog.ConsoleWriter{Out: os.Stderr})

}

func New() *cobra.Command {
	var verbosity *int = new(int)

	app := &cobra.Command{
		Use:   "gatherer",
		Short: "Gatherer",
		Long:  "Templated Terraform module gatherer",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			log.Logger = createLogger(*verbosity)
		},
		CompletionOptions: cobra.CompletionOptions{
			HiddenDefaultCmd: true,
		},
	}

	pflags := app.PersistentFlags()
	pflags.CountVarP(verbosity, "verbose", "v", "Increase verbosity level")
	app.AddCommand(NewVersionCommand())
	app.AddCommand(NewGitPullCommand())

	return app
}

func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print application version information",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(version.Get())
		},
	}
}

func NewGitPullCommand() *cobra.Command {
	var gitArgs = &git.GitArgs{}
	cmd := &cobra.Command{
		Use:   "pull",
		Short: "Pulls terraform modules from Git repository",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			fmt.Println(gitArgs)
			return nil
		},
	}

	gitArgs = git.GetGitPullArgs(cmd.Flags())

	return cmd
}
