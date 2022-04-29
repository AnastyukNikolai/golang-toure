build:
	gofmt -s -w ./
	goimports -w -d ./
	golangci-lint run ./...
	go generate ./...
	docker-compose build golang-ture

run:
	docker-compose up golang-ture

test:
	go test -v ./...

download:
	go mod download