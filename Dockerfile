FROM golang:1.18-alpine3.15 as builder

WORKDIR /app
COPY . .
RUN  CGO_ENABLED=1 CC=gcc go build -o /app/app main.go

FROM alpine:3.10

COPY --from=builder /app/app /app/
WORKDIR /app
CMD ["./app"]