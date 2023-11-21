# Run tests
`go test ./...`

# Run service
`go run ./...`

# Update swagger documentation
`export PATH=$PATH:$(go env GOPATH)/bin`  
`swag init -g ./cmd/main.go -o ./docs`

# Update schema
`go generate ./ent`

# See schema in SQL
`atlas schema inspect -u "ent://ent/schema" --dev-url "sqlite://file?mode=memory&_fk=1" --format '{{ sql . "  " }}'`
