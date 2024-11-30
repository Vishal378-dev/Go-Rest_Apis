build:
	@go build -o ./bin/app

run: build
	@echo "Running the Go application MakeFile..."
	@./bin/app
runwithpath:
	@go run main.go -config config/local.yaml