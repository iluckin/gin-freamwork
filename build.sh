#!/usr/bin/env bash

OS="darwin"

if [ "$1" == "--help" ]; then
    echo "
  Build TheBear Applicaton.
  ----------------------------
  Usage:
    - ./build.sh [os (linux | darwin)]    Compile the package for the given platform. default darwin
    - ./build.sh --help                   Show this help

  Help:
    - Get more help see : https://dev.iluckin.cn/help/bear
    "
    exit 1
fi


if [ "$1" != "" ]; then
    OS=$1
fi

echo "··> [INFO] Start compiling $OS platform application..." && GOOS="$OS" go build -o bin/bear . && echo "··> [INFO] Compile complete. Start copying configuration files ..." && cp -f cfg.toml.example bin/cfg.toml && echo "··> [INFO] OK"
