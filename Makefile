test-coverage:
	go test -v ./... -coverprofile cover.txt
	go tool cover -html=cover.txt -o cover.html
	rm -f cover.txt

test:
	go test -v ./...
lint:
	golangci-lint run -v

build:
  
