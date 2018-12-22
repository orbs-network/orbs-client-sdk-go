#!/bin/bash -xe

export DOCKER_TAG=orbs-client-sdk-go
export DOCKER_E2E_TAG=orbs-client-sdk-go-e2e
export BUILD_CONTAINER=orbs_sdk_build

docker build -f Dockerfile.build -t orbs:$DOCKER_TAG .