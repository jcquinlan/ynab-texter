build-sender:
	@echo "Building for linux"
	GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o ./cmd/sender/bootstrap ./cmd/sender
	zip ./cmd/sender/bootstrap.zip ./cmd/sender/bootstrap