.PHONY: run-dev
run-dev:
	go run cmd/ppe/ppe.go --path="$(HOME)/Projects/shark/collaboratori2/tests/Functional"

test:
	go test -v -cover -coverprofile=coverage.out ./...

view-coverage:
	go tool cover -html coverage.out 
