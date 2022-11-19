## Build
FROM golang:1.17.0-alpine

WORKDIR /go/src/github.com/seipan/scraping-go

COPY . .

RUN go get -u github.com/cosmtrek/air
RUN go mod download

CMD ["air"]