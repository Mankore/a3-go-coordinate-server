env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o go_server -ldflags="-s -w"
env GOOS=windows GOARCH=amd64 go build main.go -o go_server
