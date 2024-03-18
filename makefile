build:
	@echo "Building for linux"
	GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap .
	zip bootstrap.zip bootstrap