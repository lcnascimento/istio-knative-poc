# STAGE 0: Build binaries
FROM golang:1.14-stretch as builder

ENV CGO_ENABLED=0
ENV GOOS=linux

COPY go-libs/ /go/src/github.com/lcnascimento/istio-knative-poc/go-libs
COPY segments-service/ /go/src/github.com/lcnascimento/istio-knative-poc/segments-service

WORKDIR /go/src/github.com/lcnascimento/istio-knative-poc/segments-service
RUN rm go.mod
RUN rm go.sum

WORKDIR /go/src/github.com/lcnascimento/istio-knative-poc/segments-service/cmd/grpc/frontend
RUN go get -v
RUN go build

# STAGE 1: Build Frontend
FROM alpine as frontend

WORKDIR /go/bin

COPY --from=builder /go/src/github.com/lcnascimento/istio-knative-poc/segments-service/cmd/grpc/frontend/frontend /go/bin/frontend

RUN apk add -U --no-cache ca-certificates

ENTRYPOINT /go/bin/frontend
