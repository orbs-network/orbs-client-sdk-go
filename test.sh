#!/bin/sh -xe

go test ./crypto/... -v

go test ./test/... -v

go test ./...