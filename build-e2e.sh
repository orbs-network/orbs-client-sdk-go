#!/bin/sh +xe

CGO_ENABLED=0 time go test -ldflags '-w -extldflags "-static"' -o _bin/e2e.test -a -c ./test/e2e
