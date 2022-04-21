build:
	docker-compose build golang-ture

run:
	docker-compose up golang-ture

test:
	go test -v ./...


