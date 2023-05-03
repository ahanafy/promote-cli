package cmd

import (
	"github.com/ahanafy/promote-cli/internal/machinery"
	"github.com/spf13/cobra"
)

// promoteCmd represents the add command.
var promoteCmd = &cobra.Command{
	Use:   "promote",
	Short: "The 'promote' subcommand checks if it is safe to promote to an environment.",
	Long: `The 'promote' subcommand checks if it is safe to promote to an environment.

'<cmd> promote' will check if the environment you are promoting to is ahead of the environment you are promoting from.`,
	Run: func(cmd *cobra.Command, args []string) {
		check, _ := cmd.Flags().GetString("check")
		machinery.Start(check)
	},
}

func init() {
	rootCmd.AddCommand(promoteCmd)
	promoteCmd.PersistentFlags().StringP("check", "c", "", "Environment to check if promote-cli (required)")
}
