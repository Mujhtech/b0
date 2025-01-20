package main

import (
	"log"
	"os"

	"github.com/mujhtech/b0/cmd/hooks"
	"github.com/mujhtech/b0/cmd/migrate"
	"github.com/mujhtech/b0/cmd/server"
	"github.com/mujhtech/b0/cmd/version"
	"github.com/spf13/cobra"
)

func main() {
	err := os.Setenv("TZ", "")
	if err != nil {
		log.Fatal(err)
	}

	cmd := &cobra.Command{
		Use:     "b0",
		Version: version.Version,
		Short:   "Your AI backend builder",
	}

	// regiser hooks
	cmd.PersistentPreRunE = hooks.PreHook()
	cmd.PersistentPostRunE = hooks.PostHook()

	// Add subcommands
	cmd.AddCommand(server.RegisterServerCommand())
	cmd.AddCommand(version.RegisterVersionCommand())
	cmd.AddCommand(migrate.RegisterMigrateCommand())

	err = cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
