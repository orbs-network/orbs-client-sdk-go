#!/usr/bin/env bash
rm -rf ./_bin

mkdir -p ./_bin
go build -o _bin/gamma-cli ./gammacli

# mac only for now
tar -zcvf ./_bin/gammacli-mac-v1.2.3.tar.gz ./_bin/gamma-cli
openssl sha256 ./_bin/gammacli-mac-v1.2.3.tar.gz