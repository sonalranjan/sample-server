#
# This file was copied and modifed from
# https://github.com/GoogleCloudPlatform/spark-on-k8s-operator/blob/master/Dockerfile
#

# Using buster (for debian linux)
# https://hub.docker.com/_/golang/tags?page=1&name=buster
FROM golang:1.20.5-buster as builder

USER root
RUN apt update; apt -y install build-essential
#RUN apk update; apk add --no-cache --virtual .build-deps \
#                         build-base \
#                         curl \
#                         gcc \
#                     && apk add --no-cache \
#                         gd \
#                         libgcc

WORKDIR /workspace

# Copy the go source code
COPY main.go main.go
COPY docs/ docs/
COPY client/ client/
COPY config/ config/
COPY config/yaml /etc/sample-server-config/
COPY handler/ handler/
COPY server/ server/
COPY service/ service/

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

# Build (on mac-os arm64; use GOARCH=x for specific target architecture) 
RUN CGO_ENABLED=1 GOOS=linux GO111MODULE=on go build -a -o /usr/bin/sample-server main.go

WORKDIR /usr/bin

ENTRYPOINT ["/usr/bin/sample-server"]
