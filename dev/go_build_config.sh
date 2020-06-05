#!/bin/bash

# Build configuration options

CUR_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
SRC_ROOT="$(realpath "${CUR_DIR}/..")"

MANIFESTS_ROOT="${SRC_ROOT}/deploy"
PKG_ROOT="${SRC_ROOT}/pkg"
CONFIG_DIR="${SRC_ROOT}/config"

REPO="github.com/binarly-io/atlas-client"

# Name of the project us in all software ID and config vars
PROJECT_NAME="atlas-client"
# 1.2.3
VERSION=$(cd "${SRC_ROOT}"; cat release)
# 885c3f7
GIT_SHA=$(cd "${CUR_DIR}"; git rev-parse --short HEAD)
# 2020-01-02 12:34:56
NOW=$(date "+%FT%T")

# Check with:
# go tool dist list
# ----------------
# GOOS=
# linux
# windows
# android
# darwin
# GOARCH=
# 386
# amd64
# arm
# arm64

GOOS="${GOOS:-}"
CLIENT_BIN="${CLIENT_BIN:-}"
CGO_ENABLED="${CGO_ENABLED:-}"

# What OS are we going to build for?
if [[ -z "${GOOS}" ]]; then
    # No target OS explicitly specified, check what we have around
    if [[ "${OSTYPE}" == "cygwin" ]]; then
        # cygwin
        GOOS="windows"
    elif [[ "${OSTYPE}" == "msys" ]]; then
        # mingw
        GOOS="windows"
    elif [[ "${OSTYPE}" == "win32" ]]; then
        # windows, not sure about this
        GOOS="windows"
    else
        GOOS="linux"
    fi
fi

# Client binary name can be specified externally
if [[ "${GOOS}" == "windows" ]]; then
    if [[ -z "${CLIENT_BIN}" ]]; then
        CLIENT_BIN="${SRC_ROOT}/dev/bin/client.exe"
    fi
else
    if [[ -z "${CLIENT_BIN}" ]]; then
        CLIENT_BIN="${SRC_ROOT}/dev/bin/client"
    fi
fi

if [[ -z "${CGO_ENABLED}" ]]; then
    # Default value for CGO_ENABLED
    # CGO_ENABLED=1 means "dynamically linked"
    # CGO_ENABLED=0 means "statically linked"
    CGO_ENABLED="0"
fi

CLIENT_BUILDER_SCRIPT="go_build_client.sh"

# Where modules kept
MODULES_DIR=vendor
