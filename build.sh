#!/usr/bin/env bash

set -ex

PROJECT_DIR=$(dirname $0)
BUILD_DIR="${PROJECT_DIR}/out"

updateDependencies(){
  go get -u ./...
  go mod tidy
}

buildBinaries(){
  [[ ! -d "${BUILD_DIR}" ]] && mkdir "${BUILD_DIR}"

  set -ex

  case "$1" in
    linux)
        GOOS=linux go build -o "${BUILD_DIR}/imdb-linux-amd64" main.go
        ;;
    mac)
        GOOS=darwin go build -o "${BUILD_DIR}/imdb-darwin-amd64" main.go
        ;;
    windows)
        GOOS=windows go build -o "${BUILD_DIR}/imdb-windows-amd64" main.go
        ;;
    all)
        GOOS=linux go build -o "${BUILD_DIR}/imdb-linux-amd64" main.go
        GOOS=darwin go build -o "${BUILD_DIR}/imdb-darwin-amd64" main.go
        GOOS=windows go build -o "${BUILD_DIR}/imdb-windows-amd64" main.go
        ;;
    **)
        go build -o "${BUILD_DIR}/imdb" main.go
        ;;
  esac

  set +ex
}

##### main

case "$1" in
    -h|--help)
        set +ex
        echo "Usage: $0                   # To Build a Binary"
        echo "Usage: $0 -up-dep           # To Update Dependencies"
        echo "Usage: $0 -build            # To Build a Binary"
        echo "Usage: $0 -build all        # To Build Binary for Linux, Windoows and Mac"
        echo "Usage: $0 -build linux      # To Build Binary for Linux"
        echo "Usage: $0 -build windows    # To Build Binary for Windoows"
        echo "Usage: $0 -build mac        # To Build Binary for Mac"
        echo "Usage: $0 [-h|--help]       # For this help message"
        echo "Without a flag, defaults to build binaries."
        exit 0
        ;;
    -up-dep)
        updateDependencies
        exit 0
        ;;
    -build)
        buildBinaries $2
        exit 0
        ;;
    **)
        buildBinaries $1
        exit 0
        ;;
esac
