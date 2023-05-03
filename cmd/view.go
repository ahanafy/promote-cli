package cmd

import (
	"github.com/ahanafy/promote-cli/internal/machinery"
	"github.com/spf13/cobra"
)

// viewCmd represents the add command.
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "The 'view' subcommand checks if it is safe to view to an environment.",
	Long: `The 'view' subcommand checks if it is safe to view to an environment.

'<cmd> view' will check if the environment you are promoting to is ahead of the environment you are promoting from.`,
	Run: func(cmd *cobra.Command, args []string) {
		machinery.View()
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
