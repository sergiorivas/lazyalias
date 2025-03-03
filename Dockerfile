FROM golang:1.24-alpine AS builder

RUN apk add --no-cache \
    git \
    make \
    build-base \
    binutils \
    gcc \
    musl-dev \
    linux-headers

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.56.2

RUN make build

FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata

RUN adduser -D appuser

WORKDIR /app

COPY --from=builder /app/bin/lazyalias .

RUN chown -R appuser:appuser /app

USER appuser

ENTRYPOINT ["./lazyalias"]
