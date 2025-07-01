FROM golang:1.24.4-bullseye AS build

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o out/feed-reader cmd/main.go

FROM golang:1.24.4-alpine3.21

LABEL org.opencontainers.image.source https://github.com/marcpires/go-workshop/feed-reader
LABEL org.opencontainers.image.description "A simple feed reader for the Go workshop at LHC"

WORKDIR /app

COPY --from=build /build/out/feed-reader .

USER feed-reader

CMD ["/app/feed-ready"]
