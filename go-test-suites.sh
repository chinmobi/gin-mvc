#!/usr/bin/env bash

cd $(dirname $0)

GO_TEST_SUITES=$(find ./tests -name "*_test.go" | xargs -I {} dirname {} | sort | uniq)

echo "${GO_TEST_SUITES}"
