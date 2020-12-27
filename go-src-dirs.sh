#!/usr/bin/env bash

cd $(dirname $0)

GO_SRC_DIRS=$(find . -name "*.go" -not -path "./vendor/*" | xargs -I {} dirname {}  | uniq)

echo "${GO_SRC_DIRS}"
