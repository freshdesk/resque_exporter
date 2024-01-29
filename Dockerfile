# Build Stage
FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN cd /app/cmd/resque_exporter && go build -o /app/resque_exporter

# Final Stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/resque_exporter ./resque_exporter
COPY sample_config.yml .

EXPOSE 5555
CMD ["./resque_exporter", "--config", "sample_config.yml"]