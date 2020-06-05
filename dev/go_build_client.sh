#!/bin/bash

# Build platform-specific binary
# Do not forget to update version

# Source configuration
CUR_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
source "${CUR_DIR}/go_build_config.sh"




# Prepare modules
GO111MODULE=on go mod tidy
GO111MODULE=on go mod "${MODULES_DIR}"

if [[ -z "${OUTPUT_BIN}" ]]; then
    # Not specified, requires default value
    OUTPUT_BIN="${CLIENT_BIN}"
fi

if [[ -z "${MAIN_SRC_FILE}" ]]; then
    # Not specified, requires default value
    MAIN_SRC_FILE="${SRC_ROOT}/cmd/client/main.go"
fi

echo "Build params:"
echo "CGO_ENABLED=${CGO_ENABLED}"
echo "GOOS=${GOOS}"
echo "GOARCH=${GOARCH}"

if CGO_ENABLED="${CGO_ENABLED}" GO111MODULE=on GOOS="${GOOS}" GOARCH="${GOARCH}" go build \
    -mod="${MODULES_DIR}" \
    -a \
    -ldflags " \
        -X ${REPO}/pkg/softwareid.Name=${PROJECT_NAME} \
        -X ${REPO}/pkg/softwareid.Version=${VERSION}   \
        -X ${REPO}/pkg/softwareid.GitSHA=${GIT_SHA}    \
        -X ${REPO}/pkg/softwareid.BuiltAt=${NOW}       \
    " \
    -o "${OUTPUT_BIN}" \
    "${MAIN_SRC_FILE}"
then
    echo "Build OK"
else
    echo "WARNING! BUILD FAILED!"
    echo "Check logs for details"
    exit 1
fi
