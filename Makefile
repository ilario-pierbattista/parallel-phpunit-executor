.PHONY: build test view-coverage
build:
	go build -v ./...

test:
	go test -v -cover -coverprofile=coverage.out ./...

view-coverage:
	go tool cover -html coverage.out 

depgraph:
	godepgraph ./cmd/ppe/ | dot -Tpdf -o depgraph.pdf
