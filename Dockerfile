FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /url-shortener ./cmd/main.go

FROM alpine:3.22

RUN addgroup -S app && adduser -S app -G app

WORKDIR /app
COPY --from=builder /url-shortener /usr/local/bin/url-shortener

EXPOSE 8080
USER app

ENTRYPOINT ["/usr/local/bin/url-shortener"]
