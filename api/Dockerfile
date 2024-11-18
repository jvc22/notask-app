FROM golang:alpine AS builder

RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

ENV CGO_ENABLED=1

RUN go build -o notask-app .

FROM alpine:latest

WORKDIR /app

RUN mkdir -p /app/database/volume

COPY --from=builder /app/notask-app ./notask-app

RUN touch /app/database/volume/tasks.db

CMD ["./notask-app"]
