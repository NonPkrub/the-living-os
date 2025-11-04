FROM golang:1.25.3-alpine AS builder 
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o url-shortener .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/url-shortener .
EXPOSE 3000
CMD ["./url-shortener"]

