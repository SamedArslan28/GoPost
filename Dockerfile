FROM golang:1.25-alpine AS builder

WORKDIR /app

# For cache optimization because packages dosent change often as source code
COPY go.mod go.sum ./

run go mod download

COPY . .

WORKDIR /app/cmd/api

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/GoPost .

FROM alpine:3.18

RUN adduser -D -g '' appuser

WORKDIR /app

COPY --from=builder /app/bin/GoPost .

EXPOSE 3000

USER appuser

CMD ["./GoPost"]