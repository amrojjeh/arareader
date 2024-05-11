.PHONY: all
all: nahwapp aratools cli

.PHONY: nahwapp
nahwapp:
	go build -o ./bin/nahwapp ./cmd/nahwapp

.PHONY: aratools
aratools:
	go build -o ./bin/aratools ./cmd/aratools

.PHONY: bro
bro:
	go build -o ./bin/bro ./cmd/bro
