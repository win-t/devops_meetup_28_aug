FROM golang:alpine AS builder

RUN set -eux \
 && apk -U add git

WORKDIR /go/src/tmp/counter-app
COPY . .

RUN set -eux \
 && mkdir /output \
 && go get ./... \
 && go build -o /output/counter-app



FROM alpine:latest

RUN set -eux \
 && mkdir /app

COPY --from=builder /output/counter-app /app

ENTRYPOINT ["/app/counter-app"]
