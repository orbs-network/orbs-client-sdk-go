#!/bin/bash -xe

docker tag orbs:orbs-client-sdk-go 506367651493.dkr.ecr.us-west-2.amazonaws.com/orbs-gamma-cli
docker push 506367651493.dkr.ecr.us-west-2.amazonaws.com/orbs-gamma-cli