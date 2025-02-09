package cmd

import (
	"github.com/spf13/cobra"
)

var (
	chdir    string
	fileName string
)

func Root() *cobra.Command {
	cmd := cobra.Command{
		Use:           "runme",
		Short:         "Execute commands directly from a README",
		Long:          "Parses commands directly from a README (best-effort) to make them executable under a unique name.",
		SilenceErrors: true,
		SilenceUsage:  true,
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return nil, cobra.ShellCompDirectiveNoFileComp
		},
	}

	pflags := cmd.PersistentFlags()

	pflags.StringVar(&chdir, "chdir", ".", "Switch to a different working directory before exeucing the command.")
	pflags.StringVar(&fileName, "filename", "README.md", "A name of the README file.")

	cmd.AddCommand(runCmd())
	cmd.AddCommand(listCmd())
	cmd.AddCommand(printCmd())
	cmd.AddCommand(tasksCmd())
	cmd.AddCommand(jsonCmd())

	return &cmd
}
