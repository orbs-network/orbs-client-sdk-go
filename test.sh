#!/bin/sh -xe

go test ./crypto/... -v

go test ./gammacli/... -v

go test ./test/... -v