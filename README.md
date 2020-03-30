# euclides-api
## TL;DR How do I run this?
`./buildAndRun.sh` will create AND RUN a nice docker container for you.

Default port is `8080`.
## Configuration
- Configuration is 12-factor(-ish) via env variables. Sane defaults are provided
- `DEBUG` set to `TRUE` for debug mode - any other value or lack thereof means
that the server will be running in performance mode, with less debugging output
- `ADDRESS` which address the server should bind to
- `PORT` which port should the server listen to
- `RDTIMEOUT` Read Timeout in seconds
- `WRTIMEOUT` Write Timeout in seconds
- `IDTIMEOUT` Idle Timeout in seconds (think Slowloris ...)

## TODO
- debug/performance mode (WiP)
- optimized Ackerman
- Ackerman TEST
- Dockerize
- proper README.md
- status endpoint
## Bugs
code duplication in handlers/tests
