test:
    go test -v ./...

test-coverage:
    mkdir -p coverage

    go test -coverprofile=coverage.out.tmp ./...
    cat coverage.out.tmp | grep -v "main.go" > coverage/coverage.out
    rm coverage.out.tmp
    go tool cover -html=coverage/coverage.out
