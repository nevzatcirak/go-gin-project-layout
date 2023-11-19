FROM golang:1.21.1-alpine

# These are required by using SQLite
ENV CGO_ENABLED=1
RUN apk add build-base

#Work directory creation
RUN mkdir /app
ADD . /app
WORKDIR /app

#Building
RUN go build -o nev cmd/main.go

CMD ["/app/nev"]