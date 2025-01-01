package main

import (
	"log"
	"os"

	"github.com/mujhtech/b0/cmd/version"
	"github.com/spf13/cobra"
)

func main() {
	err := os.Setenv("TZ", "")
	if err != nil {
		log.Fatal(err)
	}

	cmd := &cobra.Command{
		Use:   "b0",
		Short: "Your AI backend builder",
	}

	// TODO: implement
	cmd.AddCommand(version.RegisterVersionCommand())

	err = cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
