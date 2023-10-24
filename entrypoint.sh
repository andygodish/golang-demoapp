#!/bin/sh
set -e

if [ "$SSH_ENABLED" = "true" ]; then
    service ssh start
fi 

./app