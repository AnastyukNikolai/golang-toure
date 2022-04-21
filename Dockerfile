FROM golang:1.16-bullseye

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update && \
    apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o golang-ture ./cmd/main.go

CMD ["./golang-ture"]