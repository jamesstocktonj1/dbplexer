package command

import "github.com/spf13/cobra"

func RootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "dbplexer",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	cmd.AddCommand(generateCommand())
	return cmd
}
