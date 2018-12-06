#!/usr/bin/env bash

python gammacli/test/generate_config.py > gammacli/test/orbs-gamma-config.json
docker-compose -f ./docker-compose.test.yml up --abort-on-container-exit --exit-code-from sdk-e2e