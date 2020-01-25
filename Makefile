check:
	go fmt ./...
	go vet ./...
	golangci-lint run

test:
	go test -p 1 ./... -tags "unit" -cover -race -a -count=1
