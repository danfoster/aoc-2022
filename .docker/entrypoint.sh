#!/bin/bash

set -e

if [ -f "cmd/day$1/main.go" ]; then
    if [ ! -f "bin/day$1" ] || [  "cmd/day$1/main.go" -nt "bin/day$1" ]; then
        export HOME="/tmp"
        go get -d -v ./... >/dev/null 2>&1
        go build -o bin/day$1 cmd/day$1/main.go >/dev/null
    fi
    export GOGC=off
    time "bin/day$1" "inputs/day$1.txt"
fi