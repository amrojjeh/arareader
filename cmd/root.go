package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	dsn string
)

var rootCmd = &cobra.Command{
	Use:   "arareader",
	Short: "A CLI which launches the best quizzing platform for Arabic",
	Long:  `A CLI which launches the best quizzing platform for Arabic. It can also be used to generate the schema and update the database.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&dsn, "dsn", "default.db", "Data Source Name. Location of the database.")
}
