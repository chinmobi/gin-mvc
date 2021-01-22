#!/usr/bin/env bash

rabbitmqctl set_policy ha-all "^" '{"ha-mode":"all"}'
