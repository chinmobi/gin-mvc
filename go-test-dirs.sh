#!/usr/bin/env bash

cd $(dirname $0)

GO_TEST_DIRS=$(find . -name "*_test.go" -not -path "./vendor/*" | xargs -I {} dirname {}  | uniq)

echo "${GO_TEST_DIRS}"
