CGO_ENABLED=1 air

go test ./... -coverprofile=coverage.out && ./exclude-from-code-coverage.sh && go tool cover -html=coverage.out