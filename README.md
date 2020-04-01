# euclides-api
## TL;DR How do I run this?
`./buildAndRun.sh` will create AND RUN a nice docker container for you.
`curl localhost:8080/help` will give you a few pointers to start
Default port is `8080`.


If you do not want/do not have Docker, `go build -o euclides && DEBUG=true ./euclides`
should do the trick

## Configuration
- Configuration is 12-factor(-ish) via env variables. Sane defaults are provided
- `DEBUG` - set to `TRUE` for debug mode - any other value or lack thereof means
that the server will be running in performance mode, with less debugging output
- `ADDRESS` - which address the server should bind to
- `PORT` - which port should the server listen to
- `RDTIMEOUT` - Read Timeout in seconds
- `WRTIMEOUT` - Write Timeout in seconds
- `IDTIMEOUT` - Idle Timeout in seconds (think Slowloris ...)

## TODO
- Ackermann TEST
- debug endpoint
## Bugs
code duplication in handlers/tests
## Not implemented
- Optimized Ackermann
