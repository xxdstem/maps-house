FROM golang:alpine AS builder

RUN apk add -U --no-cache ca-certificates

ARG TARGETPLATFORM
ARG TARGETOS 
ARG TARGETARCH

WORKDIR /build

ADD go.mod .

COPY . .

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o migrate cmd/migrate/migrate.go

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o app cmd/app/main.go

FROM --platform=${TARGETPLATFORM:-linux/amd64} alpine

WORKDIR /

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app

COPY --from=builder /build/migrate migrate
COPY --from=builder /build/app app
COPY --from=builder /build/app.env app.env

COPY --from=builder /build/start.sh /app/start.sh

RUN chmod +x /app/start.sh

ENTRYPOINT ["/bin/sh", "/app/start.sh"]