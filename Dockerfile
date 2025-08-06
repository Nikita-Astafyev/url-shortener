FROM golang:1.24.0-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o url-shortener ./cmd/server

# Финальный образ
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/url-shortener .

RUN mkdir -p configs && \
    (cp -r /app/configs/. ./configs/ 2>/dev/null || true)

EXPOSE 8080
CMD ["./url-shortener"]