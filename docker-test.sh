#!/usr/bin/env bash

rm -rf _tmp
mkdir _tmp

export GAMMA_ENDPOINT=$(python gammacli/test/generate_config.py endpoint)
python gammacli/test/generate_config.py json > _tmp/orbs-gamma-config.json

docker-compose -f ./docker-compose.test.yml up --abort-on-container-exit --exit-code-from sdk-e2e