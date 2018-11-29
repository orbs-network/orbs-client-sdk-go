rm -rf ./_bin

mkdir -p ./_bin
go build -o ./_bin/gamma-cli

# mac only for now
tar -zcvf ./_bin/gammacli-mac.tar.gz ./_bin/gamma-cli
openssl sha256 ./_bin/gammacli-mac.tar.gz