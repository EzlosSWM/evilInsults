build: 
	@go build -o bin/evilInsults 

run: build
	@./bin/evilInsults

test: 
	go test -v ./...