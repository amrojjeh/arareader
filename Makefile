.PHONY: coverage
coverage:
	go test ./... -covermode=count -coverprofile=c.out
	go tool cover -html=c.out
