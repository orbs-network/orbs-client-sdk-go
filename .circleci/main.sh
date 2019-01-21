#!/bin/bash -xe

# First let's install Go 1.11

echo "Installing Go 1.11"
cd /tmp
wget https://dl.google.com/go/go1.11.linux-amd64.tar.gz
sudo tar -xvf go1.11.linux-amd64.tar.gz
# Uninstall older version of Go
sudo rm -rf /usr/local/go
sudo mv go /usr/local

export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

go version

./.circleci/bring-gamma.sh
./test.sh