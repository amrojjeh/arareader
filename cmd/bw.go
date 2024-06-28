package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/amrojjeh/arareader/arabic"
	"github.com/spf13/cobra"
)

// bwCmd represents the bw command
var bwCmd = &cobra.Command{
	Use:     "bw",
	Short:   "Go from either Arabic to Buckwalter or from Buckwalter to Arabic",
	Long:    `Go from either Arabic to Buckwalter or from Buckwalter to Arabic`,
	Example: `arareader bw "h*A baytN"`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("hosting http://localhost:8080")
		http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet {
				err := r.ParseForm()
				if err != nil {
					panic("parsing form")
				}
				ar := r.Form.Get("ar")
				bw := arabic.ToBuckwalter(ar)
				w.Write([]byte(fmt.Sprintf(`
<!doctype html>
<html>
	<body>
		<form action="/" method="get">
			Arabic
			<textarea name="ar" rows="50" cols="100" type="text" dir="rtl" autofocus>%s</textarea>
			Buckwalter
			<textarea type="text" rows="50" cols="100" readonly>%s</textarea>
			<button type="submit">Transliterate</button>
		</form>
	</body>
</html>`, ar, bw)))
			}
		}))
	},
}

func init() {
	rootCmd.AddCommand(bwCmd)
}
