FROM golang:1.25-alpine AS builder

WORKDIR /build-stage

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath -ldflags="-s -w" \
    -o /build-stage/service ./cmd/app

FROM alpine:3.22

RUN apk add --no-cache ca-certificates tzdata

RUN adduser -D -u 10001 appuser

WORKDIR /app

COPY --from=builder /build-stage/service /app/service

RUN ls -l /app/service
RUN chmod +x /app/service
RUN ls -l /app/service
RUN file /app/service

USER appuser

ENTRYPOINT [ "/app/service" ]