FROM golang:1.25.4-alpine AS builder

WORKDIR /app

ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/server/main.go


FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/server ./server
COPY --from=builder /app/config ./config
COPY --from=builder /app/api ./api

EXPOSE 8080

CMD ["./server"]