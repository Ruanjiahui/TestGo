#!/usr/bin/env bash

CURRENT_DIR=`pwd`
OLD_GO_PATH="Z:\Go work"  #例如: /usr/local/go
OLD_GO_BIN="Z:\Go\bin"    #例如: /usr/local/go/bin

export GOPATH="$CURRENT_DIR" 
export GOBIN="$CURRENT_DIR/bin"

#指定并整理当前的源码路径
gofmt -w src

go install main

export GOPATH="$OLD_GO_PATH"
export GOBIN="$OLD_GO_BIN"