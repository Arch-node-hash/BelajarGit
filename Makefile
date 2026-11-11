.PHONY: run test vet fmt

run:
	go run . list

test:
	go test ./...

vet:
	go vet ./...

fmt:
	gofmt -w .
