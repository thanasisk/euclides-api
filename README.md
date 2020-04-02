# euclides-api
## TL;DR How do I run this?
`./buildAndRun.sh` will create AND RUN a nice docker container for you.


`curl localhost:8080/help` will give you a few pointers to start


Default port is `8080`.


If you do not want/do not have Docker, `go build -o euclides && DEBUG=true ./euclides`
should do the trick. Remember to run `file(1)` on the binary to dispell a common Golang misconception :-)
## Read the source! Especially comments!
See above
## Some security measures
- We do not run as root
- SSL is expected to be handled by a reverse proxy (think nginx-ingress or similar)
- Differentiation between performance and debug mode
## Testing
`go test -v` will perform unit testing
## Configuration
- Configuration is 12-factor(-ish) via env variables. Sane defaults are provided
- `DEBUG` - set to `TRUE` for debug mode - any other value or lack thereof means
that the server will be running in performance mode, with less debugging output
- `ADDRESS` - which address the server should bind to
- `PORT` - which port should the server listen to
- `RDTIMEOUT` - Read Timeout in seconds
- `WRTIMEOUT` - Write Timeout in seconds
- `IDTIMEOUT` - Idle Timeout in seconds (think Slowloris ...)

## Bugs
some code duplication

## Not implemented
- Optimized Ackermann

## 3rd party software
contains code by the gorilla/mux team

## LICENSE
GPL of course :D
