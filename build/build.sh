#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

PACKAGE="github.com/openshift/windows-machine-config-operator"
MAIN_PACKAGE="${PACKAGE}/cmd/manager"
BIN_NAME="windows-machine-config-operator"
OUTPUT_DIR="build/_output"
BIN_DIR="${OUTPUT_DIR}/bin"

echo "building ${BIN_NAME}..."
mkdir -p "${BIN_DIR}"

# Account for environments where GOFLAGS is not set by setting goflags to an empty string if GOFLAGS is not set
goflags=${GOFLAGS:-}

# The golang 1.13 image used in CI enforces vendoring. Workaround that by unsetting it.
if [[ "$goflags" == *"-mod=vendor"* ]]; then
  unset GOFLAGS
fi
CGO_ENABLED=0 GO111MODULE=on GOOS=linux go build -o ${BIN_DIR}/${BIN_NAME} ${MAIN_PACKAGE}