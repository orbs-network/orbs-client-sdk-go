#!/usr/bin/env bash

docker-compose -f ./docker-compose.test.yml up --abort-on-container-exit --exit-code-from sdk-e2e