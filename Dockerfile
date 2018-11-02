FROM golang:alpine

RUN apk update && \
    apk upgrade && \
    apk add git