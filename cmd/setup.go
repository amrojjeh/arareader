/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package cmd

import (
	"github.com/amrojjeh/arareader/service"
	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Primes the application ready for use",
	Long:  `Sets up the database by creating the schema`,
	Run: func(cmd *cobra.Command, args []string) {
		db := service.OpenDB(dsn)
		service.Setup(cmd.Context(), db)
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
