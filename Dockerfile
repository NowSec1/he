FROM golang:alpine as builder

RUN apk add --no-cache make git
WORKDIR /he-src
COPY . /he-src
RUN go mod download && \
    make linux-amd64 && \
    mv ./bin/he-linux-amd64 /he

FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /he /
ENTRYPOINT ["/he"]