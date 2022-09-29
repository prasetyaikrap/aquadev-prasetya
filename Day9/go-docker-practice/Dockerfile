FROM golang:alpine AS builder

#build
COPY . /app
WORKDIR /app
RUN apk update \
    && apk upgrade \
    && apk add --no-cache \
    ca-certificates \
    && update-ca-certificates 2>/dev/null || true
RUN mkdir app && \
    mkdir app/resources && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /app/echo
RUN ls app

# SERVE
FROM busybox

COPY --from=builder /app/ /app
WORKDIR /app
EXPOSE 1323
CMD ["./echo"]