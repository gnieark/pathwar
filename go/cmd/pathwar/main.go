package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/peterbourgon/ff"
	"github.com/peterbourgon/ff/ffcli"
	"pathwar.land/pathwar/v2/go/pkg/pwversion"
)

func main() {
	log.SetFlags(0)

	defer func() {
		if logger != nil {
			_ = logger.Sync()
		}
	}()

	// setup flags
	var (
		globalFlags = flag.NewFlagSet("pathwar", flag.ExitOnError)
	)
	globalFlags.SetOutput(flagOutput) // used in main_test.go
	globalFlags.BoolVar(&globalDebug, "debug", false, "debug mode")
	globalFlags.StringVar(&zipkinEndpoint, "zipkin-endpoint", "", "optional opentracing server")
	globalFlags.StringVar(&bearerSecretKey, "bearer-secretkey", "", "bearer.sh secret key")
	globalFlags.StringVar(&globalSentryDSN, "sentry-dsn", "", "Sentry DSN")

	version := &ffcli.Command{
		Name:      "version",
		Usage:     "pathwar [global flags] version",
		ShortHelp: "show version",
		Exec: func(args []string) error {
			fmt.Printf(
				"version=%q\ncommit=%q\nbuilt-at=%q\nbuilt-by=%q\n",
				pwversion.Version, pwversion.Commit, pwversion.Date, pwversion.BuiltBy,
			)
			return nil
		},
	}

	root := &ffcli.Command{
		Usage:    "pathwar [global flags] <subcommand> [flags] [args...]",
		FlagSet:  globalFlags,
		LongHelp: "More info here: https://github.com/pathwar/pathwar/wiki/CLI",
		Subcommands: []*ffcli.Command{
			rawclientCommand(),
			apiCommand(),
			composeCommand(),
			agentCommand(),
			miscCommand(),
			adminCommand(),
			version,
		},
		Options: []ff.Option{ff.WithEnvVarNoPrefix()},
		Exec:    func([]string) error { return flag.ErrHelp },
	}

	if err := root.Run(os.Args[1:]); err != nil {
		if !errors.Is(err, flag.ErrHelp) {
			log.Fatalf("fatal: %+v", err)
		}
	}
}
