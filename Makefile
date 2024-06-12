.PHONY: coverage
coverage:
	go test ./... -covermode=count -coverprofile=c.out
	go tool cover -html=c.out

.PHONE: sass
sass:
	sass main.scss ui/static/main.css --style=compressed
