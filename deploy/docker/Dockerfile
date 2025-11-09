FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o backend ./cmd


FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/backend .
COPY --from=builder /app/config ./config
EXPOSE 8082
CMD ["./backend"]
