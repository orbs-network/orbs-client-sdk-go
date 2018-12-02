#!/bin/sh +xe

CGO_ENABLED=0 time go test -o _bin/e2e.test -a -c ./test/e2e
