build:
	docker-compose build golang-ture

run:
	docker-compose up golang-ture

generate:
	go generate ./ent

test:
	go test -v ./...

download:
	go mod download