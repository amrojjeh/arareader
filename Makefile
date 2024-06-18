.PHONY: coverage
coverage:
	go test ./... -covermode=count -coverprofile=c.out
	go tool cover -html=c.out

.PHONY: demo
demo: tailwind
	go run . start demo

.PHONY: tailwind
tailwind:
	./bin/tailwindcss -i input.css -o ./ui/static/main.css
