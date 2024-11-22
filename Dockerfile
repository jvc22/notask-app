FROM golang:alpine AS compiler

RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

COPY ./api/go.mod ./api/go.sum ./

RUN CGO_ENABLED=0 go mod tidy -e && \
    CGO_ENABLED=0 go mod download

RUN go mod download

COPY ./api/ ./

ENV CGO_ENABLED=1

RUN go build -o api .

FROM node:20.18.0-alpine3.19 AS base

FROM base AS deps

RUN apk add --no-cache libc6-compat

WORKDIR /app

COPY ./client/package*.json ./

RUN npm ci --force

FROM base AS builder

WORKDIR /app

COPY --from=deps /app/node_modules ./node_modules

COPY ./client ./

RUN npm run build

RUN npm prune --production --force

FROM base AS runner

WORKDIR /app

RUN addgroup --system --gid 1001 nodejs

RUN adduser --system --uid 1001 nextjs

COPY --from=builder --chown=nextjs:nodejs /app/.next/standalone ./

COPY --from=builder --chown=nextjs:nodejs /app/.next/static ./.next/static

RUN mkdir -p /app/api

COPY --from=compiler /app/api ./api/api

COPY .env .

RUN mkdir -p /app/database/volume && \
    chmod 777 /app/database/volume && \
    touch /app/database/volume/tasks.db && \
    chmod 666 /app/database/volume/tasks.db

USER nextjs

EXPOSE 3000

EXPOSE 8080

ENV PORT=3000

ENV HOSTNAME="0.0.0.0"

CMD ["sh", "-c", "./api/api & node server.js"]

