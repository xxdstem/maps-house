FROM golang:alpine AS builder

RUN apk add -U --no-cache ca-certificates

ARG TARGETPLATFORM
ARG TARGETOS 
ARG TARGETARCH

WORKDIR /build

ADD go.mod .

COPY . .

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o bot cmd/bot/main.go

FROM --platform=${TARGETPLATFORM:-linux/amd64} scratch

WORKDIR /

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app

ENV TELEGRAM_KEY ""
ENV OPENAI_KEY ""
ENV LIMITER_INTERVAL 10

COPY --from=builder /build/bot bot

ENTRYPOINT ["/app/bot"]