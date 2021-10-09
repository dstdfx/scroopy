#!/usr/bin/env bash

echo "==> Building scroopy-repl binary..."
GO111MODULE=on CGO_ENABLED=0 \
go build -mod=mod -a -installsuffix cgo -ldflags \
    "-X scroopy/cmd/repl/app.buildGitCommit=$(git rev-parse HEAD) \
    -X scroopy/cmd/repl/app.buildGitTag=$(git describe --abbrev=0) \
    -X scroopy/cmd/repl/app.buildDate=$(date +%Y%m%d)" \
    -o scroopy-repl ./cmd/repl
