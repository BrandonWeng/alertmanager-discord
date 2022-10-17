#!/bin/bash
set -e
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PARENT=`dirname ${DIR}`
NAME=alertmanager-discord

VERSION=`cat ${PARENT}/VERSION`
IMAGE=ghcr.io/BrandonWeng/${NAME}:${VERSION}
