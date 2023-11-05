# Run tests
`go test ./...`

# Run service
`go run ./...`

# Update swagger documentation
`export PATH=$PATH:$(go env GOPATH)/bin`  
`swag init -g ./cmd/main.go -o ./docs`