FROM golang:alpine AS builder

RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN CGO_ENABLED=0 go mod tidy -e && \
    CGO_ENABLED=0 go mod download

RUN go mod download

COPY ./ ./

ENV CGO_ENABLED=1

RUN go build -o api .

FROM alpine:latest

WORKDIR /app

RUN mkdir -p /app/database/volume

COPY --from=builder /app/api ./api

RUN touch /app/database/volume/tasks.db

EXPOSE 8080

CMD ["./api"]
