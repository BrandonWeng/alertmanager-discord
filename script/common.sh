#!/bin/bash
set -e
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PARENT=`dirname ${DIR}`
1=alertmanager-discord

VERSION=`cat ${PARENT}/VERSION`
IMAGE=ghcr.io/brandonweng/${NAME}:${VERSION}
