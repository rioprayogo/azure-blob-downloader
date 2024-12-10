FROM golang:1.20 as builder

WORKDIR /app
COPY . .
RUN go mod tidy && go build -o /app/azure-downloader ./cmd/downloader

FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /root/
COPY --from=builder /app/azure-downloader .

ENTRYPOINT ["./azure-downloader"]
