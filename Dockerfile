FROM golang:latest

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR ./cmd/api

RUN go build -o app

EXPOSE 7777

CMD ["./app"]