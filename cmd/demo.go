/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package cmd

import (
	"log"

	"github.com/amrojjeh/arareader/routes"
	"github.com/amrojjeh/arareader/service"
	"github.com/spf13/cobra"
)

// demoCmd represents the demo command
var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "Start the server with dummy data demonstration",
	Long:  `Creates an in-memory database and fills it up with dummy data so that it could be toyed and tested without being reset each time. It's also useful for demonstrations.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := service.DemoDB(cmd.Context())
		handler := routes.NewRootHandler(db)
		server := routes.Server(handler, addr())
		log.Printf("Listening on %s...", addr())
		if err := server.ListenAndServeTLS(certPath(), keyPath()); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	startCmd.AddCommand(demoCmd)
}
