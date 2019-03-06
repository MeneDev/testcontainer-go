#!/usr/bin/env bash

STARTUP_DELAY=${STARTUP_DELAY:-"2s"}

echo "Waiting ${STARTUP_DELAY}"
sleep ${STARTUP_DELAY}
exec "$@"
