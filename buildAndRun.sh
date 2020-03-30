#!/bin/sh
set -e
docker build -t euclides .
docker run -it -p8080:8080 -eDEBUG=true euclides
