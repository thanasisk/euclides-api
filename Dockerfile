# Debian Based images
FROM golang:1.14.1 as builder
WORKDIR /go
ADD . /go
RUN go get -u github.com/gorilla/mux # scalable for only a few deps
RUN go test -v
# without -ldflags it will create a DYNAMIC binary, brrrrr
RUN go build  -ldflags "-linkmode external -extldflags -static" -o /go/euclides
# why debian-buster again? see output of go build, using alpine or scratch leads to hard errors
#/usr/bin/ld: /tmp/go-link-779577606/000004.o: in function `_cgo_26061493d47f_C2func_getaddrinfo':
#/tmp/go-build/cgo-gcc-prolog:58: warning: Using 'getaddrinfo' in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
FROM debian:buster-slim
LABEL MAINTAINER <athanasios@akostopoulos.com>
COPY --from=builder /go/euclides /
EXPOSE 8080
USER 1001
CMD /euclides
