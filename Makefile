
build:
	go build -o ./cmd/server/main.go
compile:
	GOOS=linux GOARCH=amd64 go build -v github.com/1r0npipe/shortener-web-links
	GOOS=windows GOARCH=amd64 go build -v github.com/1r0npipe/shortener-web-links
	GOOS=darwin GOARCH=amd64 go build -v github.com/1r0npipe/shortener-web-links
run:
	go run ./cmd/server/main.go --config config.yaml