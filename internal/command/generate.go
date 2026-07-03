package command

import (
	"fmt"

	"github.com/jamesstocktonj1/dbplexer/pkg/generator"
	"github.com/spf13/cobra"
)

func generateCommand() *cobra.Command {
	opts := generator.DefaultBuildOptions()
	cmd := &cobra.Command{
		Use: "generate",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("expected only 1 argument but got %d", len(args))
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.PackageName = args[0]
			return generator.NewGenerator().Run(opts)
		},
	}
	cmd.Flags().StringVarP(&opts.Output, "output", "o", "output", "output directory of generated shim")
	cmd.Flags().StringArrayVar(&opts.DatabaseNames, "names", []string{}, "names for databases to connect to")

	return cmd
}
