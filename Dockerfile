FROM golang:1.24-alpine AS build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux


WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -ldflags "-s -w" -v  -o books-api

FROM alpine:latest AS final


COPY --from=build /app/books-api /app/books-api


EXPOSE 7070

ENTRYPOINT ["/app/books-api"]
