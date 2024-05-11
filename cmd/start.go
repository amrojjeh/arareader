/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package cmd

import (
	"fmt"

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
		// TODO(Amr Ojjeh): Implement
		fmt.Println("Not yet implemented...")
		// tlsPath = path.Clean(tlsPath)
		// log.Printf("listening on %s...\n", addr())
		// server := service.Server(log.Default(), service.HTTPRoute{}, addr())
		// if err := server.ListenAndServeTLS(certPath(), keyPath()); err != nil {
		// log.Println(err)
		// }
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "specifies the port of the server")
	rootCmd.PersistentFlags().StringVar(&tlsPath, "tls", "./testtls", "a path to the cert and private key for https")
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
