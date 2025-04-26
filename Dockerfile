FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/etl ./cmd

FROM alpine:3.18 AS runner

WORKDIR /

COPY --from=builder /app/etl /etl

COPY .env /.env

CMD ["/etl"]
