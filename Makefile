dependence:
	go mod tidy
	go mod vendor

test: dependence
	go env -w CGO_ENABLED=1
	go test -cover -race ./...


.PHONY: dependence test
