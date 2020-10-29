# Build Gpch in a stock Go builder container
FROM golang:1.11-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /go-psdchaineum
RUN cd /go-psdchaineum && make gpch

# Pull Gpch into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-psdchaineum/build/bin/gpch /usr/local/bin/

EXPOSE 8545 8546 30303 30303/udp
ENTRYPOINT ["gpch"]
