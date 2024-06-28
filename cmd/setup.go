package cmd

import (
	"github.com/amrojjeh/arareader/model"
	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Primes the application ready for use",
	Long:  `Sets up the database by creating the schema`,
	Run: func(cmd *cobra.Command, args []string) {
		db := model.MustOpenDB(dsn)
		model.MustSetup(cmd.Context(), db)
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
