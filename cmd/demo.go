/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package cmd

import (
	"log"
	"net/http"
	"time"

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
		handler := routes.Routes(db)
		server := server(handler)
		log.Printf("Listening on %s...", addr())
		if err := server.ListenAndServeTLS(certPath(), keyPath()); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	startCmd.AddCommand(demoCmd)
}

func server(handler http.Handler) http.Server {
	return http.Server{
		Addr:              addr(),
		Handler:           handler,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       15 * time.Second,
		ErrorLog:          log.Default(),
	}
}
