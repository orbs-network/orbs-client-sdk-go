#!/bin/bash -xe

export DOCKER_TAG=orbs-client-sdk-go
export BUILD_CONTAINER=orbs_sdk_build

docker build -f Dockerfile -t orbs:$DOCKER_TAG .

[ "$(docker ps -a | grep $BUILD_CONTAINER)" ] && docker rm -f $BUILD_CONTAINER

docker run --name $BUILD_CONTAINER orbs:$DOCKER_TAG sleep 1

export SRC=/go/src/github.com/orbs-network/orbs-client-sdk-go

rm -rf _bin
docker cp $BUILD_CONTAINER:$SRC/_bin .
