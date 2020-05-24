FROM golang:1.14-alpine AS builder

WORKDIR /go/src/github.com/drasko/k8s-tutorial/lesson-1
COPY . .
RUN apk update \
    && apk add make \
    && make \
    && mv app /exe

FROM scratch
COPY --from=builder /exe /
ENTRYPOINT ["/exe"]
