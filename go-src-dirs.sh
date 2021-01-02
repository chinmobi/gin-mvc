#!/usr/bin/env bash

cd $(dirname $0)

GO_SRC_DIRS=$(find . -name "*.go" ! -path "./_deploy/*" | xargs -I {} dirname {} | sort | uniq)

echo "${GO_SRC_DIRS}"
