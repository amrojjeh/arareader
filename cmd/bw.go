/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package cmd

import (
	"fmt"
	"strings"

	"github.com/amrojjeh/arareader/arabic"
	"github.com/spf13/cobra"
)

var fromArabic bool

// bwCmd represents the bw command
var bwCmd = &cobra.Command{
	Use:     "bw",
	Short:   "Go from either Arabic to Buckwalter or from Buckwalter to Arabic",
	Long:    `Go from either Arabic to Buckwalter or from Buckwalter to Arabic`,
	Example: `arareader bw "h*A baytN"`,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		phrase := strings.Builder{}
		for _, w := range args {
			phrase.WriteString(w)
			phrase.WriteRune(' ')
		}
		if fromArabic {
			fmt.Println(arabic.ToBuckwalter(phrase.String()))
			return
		}
		fmt.Println(arabic.FromBuckwalter(phrase.String()))
	},
}

func init() {
	rootCmd.AddCommand(bwCmd)

	bwCmd.Flags().BoolVar(&fromArabic, "ar", false, "go from Arabic to Buckwalter")
}
