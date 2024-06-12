.PHONY: coverage
coverage:
	go test ./... -covermode=count -coverprofile=c.out
	go tool cover -html=c.out

.PHONY: sass
sass:
	sass main.scss ui/static/main.css --style=compressed

.PHONY: sass-watch
sass-watch:
	sass main.scss ui/static/main.css --style=compressed -w
