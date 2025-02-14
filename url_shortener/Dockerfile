FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM alpine:latest
# RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .

ENV POSTGRES_USER=user
ENV POSTGRES_PASSWORD=password
ENV POSTGRES_DB=db_name

COPY init.sql /docker-entrypoint-initdb.d/

ENV PORT=8080

CMD ["./main"]
