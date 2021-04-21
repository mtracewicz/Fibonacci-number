#!/bin/sh
docker build . -t gofib
docker run -p 4444:80 --rm gofib