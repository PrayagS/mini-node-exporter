FROM golang:1.17.0-buster as builder
WORKDIR /mini-node-exporter
COPY . .
RUN CGO_ENABLED=0 go build -o main ./cmd/main.go

FROM debian:buster-slim
WORKDIR /app
COPY --from=builder /mini-node-exporter/main .
CMD ["./main"]
