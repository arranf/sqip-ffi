#!/usr/bin/env bash
set -e
go build -buildmode=c-archive -o ./build/libsqip.a main.go