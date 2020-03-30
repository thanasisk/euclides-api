# Debian Based images
FROM golang:1.14.1 as builder
WORKDIR /go
ADD . /go
RUN go get -u github.com/gorilla/mux # scalable for only a few deps
RUN go test -v
# without -ldflags it will create a DYNAMIC binary, brrrrr
RUN go build  -ldflags "-linkmode external -extldflags -static" -o /go/euclides
FROM debian:buster-slim
LABEL MAINTAINER <athanasios@akostopoulos.com>
COPY --from=builder /go/euclides /
EXPOSE 8080
CMD /euclides
