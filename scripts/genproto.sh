#!/usr/bin/env bash
#
# Generate all etcd protobuf bindings.
# Run from repository root directory named etcd.
#
set -e
shopt -s globstar

if ! [[ "$0" =~ scripts/genproto.sh ]]; then
  echo "must be run from repository root"
  exit 255
fi

source ./scripts/test_lib.sh

# fork note (opaque-api branch): accept any protoc >= 3.20.3 so the hybrid
# API files can be generated with a locally installed newer protoc.
if ! [[ $(protoc --version | cut -f2 -d' ') =~ ^[0-9]+\. ]]; then
  echo "could not find protoc, is it installed + in PATH?"
  exit 255
fi

GOGEN_BIN=$(tool_get_bin google.golang.org/protobuf/cmd/protoc-gen-go)

echo
echo "Resolved binary and packages versions:"
echo "  - protoc-gen-go:           ${GOGEN_BIN}"

# directories containing protos to be built
DIRS="./raftpb"

log_callout -e "\\nRunning protoc-gen-go proto generation..."

for dir in ${DIRS}; do
  pushd "${dir}"
    protoc --go_out=. -I=".:${RAFT_ROOT_DIR}/..:${RAFT_ROOT_DIR}" \
    --plugin=protoc-gen-go="${GOGEN_BIN}" \
    --go_opt=paths=source_relative --go_opt=default_api_level=API_HYBRID ./**/*.proto

    rm -f ./**/*.bak
    gofmt -s -w ./**/*.pb.go
    run_go_tool "golang.org/x/tools/cmd/goimports" -w ./**/*.pb.go
  popd
done

log_success -e "\\n./genproto SUCCESS"
