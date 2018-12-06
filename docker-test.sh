#!/usr/bin/env bash

export GAMMA_ENDPOINT=$(python gammacli/test/generate_config.py endpoint)
python gammacli/test/generate_config.py json > gammacli/test/orbs-gamma-config.json

docker-compose -f ./docker-compose.test.yml up --abort-on-container-exit --exit-code-from sdk-e2e