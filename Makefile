gen-mock:
	mockery --dir ./database/postgresql/adapter --name Postgres  --output mocks/mock_adapter --outpkg mock_adapter

test:
	go test ./...
