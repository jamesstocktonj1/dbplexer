package command

import "github.com/spf13/cobra"

func RootCommand() *cobra.Command {
	return &cobra.Command{
		Use: "dbplexer",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
}
