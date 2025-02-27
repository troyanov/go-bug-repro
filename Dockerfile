FROM ubuntu:24.04

COPY --from=golang:1.23 /usr/local/go /usr/local/go
ENV PATH="/usr/local/go/bin:${PATH}"

RUN apt update && apt install -y \
    gcc \
    software-properties-common

RUN add-apt-repository ppa:dqlite/dev
RUN apt update && apt install -y \
    libdqlite1.17-dev \
    libdqlite1.17-0

WORKDIR /app/src

COPY main.go go.mod go.sum .
