.PHONY: ci

ci:
	make install
	make gen-mock
	make test

install:
	go install github.com/vektra/mockery/v2@latest

gen-mock:
	mockery --dir ./database/postgresql/adapter --name Postgres  --output mocks/mock_adapter --outpkg mock_adapter

test:
	go test ./...
