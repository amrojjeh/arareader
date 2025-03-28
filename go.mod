module github.com/amrojjeh/arareader

go 1.22

require (
	github.com/alexedwards/scs/sqlite3store v0.0.0-20240316134038-7e11d57e8885
	github.com/alexedwards/scs/v2 v2.8.0
	github.com/go-chi/chi/v5 v5.0.12
	github.com/maragudk/gomponents v0.20.3
	github.com/maragudk/gomponents-htmx v0.5.0
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/spf13/cobra v1.8.0
	github.com/stretchr/testify v1.9.0
	golang.org/x/crypto v0.23.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/maragudk/gomponents => ./forks/gomponents
