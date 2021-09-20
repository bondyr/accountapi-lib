FROM golang:1.17.1 AS base

FROM base AS unittest
WORKDIR /files
COPY go.mod /files
COPY go.sum /files
COPY src /files/src
WORKDIR /files/src/go/v1
RUN go test -v

FROM unittest AS libtest
COPY test /files/test
WORKDIR /files/test
CMD [ "go", "test", "-v" ]
