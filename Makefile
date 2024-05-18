build:
	go build -o bin/usafi-hub-cleaning-service
	
serve: build
	ENV=development ./bin/usafi-hub-cleaning-service

test: build
	ENV=development_test go test -v ./...