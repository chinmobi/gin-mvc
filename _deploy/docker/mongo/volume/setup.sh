#!/usr/bin/env bash

cd $(dirname $0) && source ../.env


if ! docker volume ls | grep -q "${VOLUME_NAME}"; then
  docker volume create -d local ${VOLUME_NAME}
  docker volume create -d local ${CONFIG_VOLUME_NAME}
fi

if [ ! -d "${HOME_MOUNT}" ]; then
  mkdir -p ${HOME_MOUNT} && touch ${HOME_MOUNT}/.dbshell
  sudo chmod -R a+w ${HOME_MOUNT}
fi
