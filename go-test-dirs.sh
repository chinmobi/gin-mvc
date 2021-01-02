#!/usr/bin/env bash

cd $(dirname $0)

GO_TEST_DIRS=$(find . -name "*_test.go" ! -path "./tests/*" ! -path "./_deploy/*" | xargs -I {} dirname {} | sort | uniq)

echo "${GO_TEST_DIRS}"
