FROM golang:1.16.5-alpine as builder

LABEL maintainer="Shuang Pan <panshuangis88188@@gmail.com>"

ENV GO111MODULE=on
ENV CASSANDRA_URL=cassandra:9042

RUN apk update && apk add --no-cache git

COPY apis /go/src/secretBoxAPI/apis
COPY go.mod go.sum main.go /go/src/secretBoxAPI/

WORKDIR /go/src/secretBoxAPI

RUN chmod 755 ./apis/build/build_binary.sh
RUN sh ./apis/build/build_binary.sh

FROM scratch

COPY --from=builder /go/src/secretBoxAPI/CS157C-TEAM8 .

CMD ["./CS157C-TEAM8"]
