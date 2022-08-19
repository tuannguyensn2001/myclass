package cmd

import "github.com/spf13/cobra"

func Root() *cobra.Command {
	cmdRoot := &cobra.Command{}

	commands := []*cobra.Command{server(), migrateRefresh(), migrateUp(), migrateDown(), genError(), testValidate()}

	for _, item := range commands {
		cmdRoot.AddCommand(item)
	}

	return cmdRoot
}
