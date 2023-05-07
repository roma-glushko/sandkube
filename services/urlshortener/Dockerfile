FROM golang:1.20-alpine3.17 as base

WORKDIR /service

FROM base as build

COPY go.mod go.sum main.go /service
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /service/urlshortener

FROM alpine:3.17.3 AS release

WORKDIR /service

COPY --from=build /service/urlshortener /service/urlshortener

EXPOSE 8000

ENTRYPOINT ["/service/urlshortener"]
