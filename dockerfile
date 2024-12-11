FROM golang:1.20 AS builder

WORKDIR /app
COPY . .
RUN go mod tidy && go build -o /app/azure-downloader ./cmd/downloader

ENTRYPOINT ["./azure-downloader"]