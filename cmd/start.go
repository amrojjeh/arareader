/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package cmd

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/amrojjeh/arareader/service"
	"github.com/spf13/cobra"
)

var (
	port    int
	tlsPath string
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the Arareader web server",
	Long:  `Starts the Arareader web server.`,
	Run: func(cmd *cobra.Command, args []string) {
		tlsPath = path.Clean(tlsPath)
		server := http.Server{
			Addr:              addr(),
			Handler:           service.HTTPRoute{},
			ReadTimeout:       15 * time.Second,
			ReadHeaderTimeout: 15 * time.Second,
			WriteTimeout:      15 * time.Second,
			IdleTimeout:       15 * time.Second,
			ErrorLog:          log.Default(),
		}
		log.Printf("listening on %s...\n", addr())
		if err := server.ListenAndServeTLS(certPath(), keyPath()); err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.Flags().IntVarP(&port, "port", "p", 8080, "specifies the port of the server")
	rootCmd.Flags().StringVar(&tlsPath, "tls", "./testtls", "a path to the cert and private key for https")
}

func certPath() string {
	return fmt.Sprintf("%s/fullchain.pem", tlsPath)
}

func keyPath() string {
	return fmt.Sprintf("%s/privkey.pem", tlsPath)
}

func addr() string {
	return fmt.Sprintf(":%d", port)
}
