#!/usr/bin/env bash

cd $(dirname $0) && source ./dev.env

exec ./bin/${APP_BIN_NAME}
