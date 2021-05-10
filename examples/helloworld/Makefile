# vim:set noet

# 当前的用户 `whoami` 命令
USER?=`whoami`  # ?= 来接受外部环境变量重写
GIT_TAG=`git describe --tags --always`
GIT_COMMIT=`git rev-parse HEAD`
BUILD_TIME=`date +%FT%T`
PKG_VERSION=github.com/yogorobot/marsaxlokk/version
BUILD_BUILD_FLAGS=-ldflags "-X ${PKG_VERSION}.Version=${GIT_TAG} -X ${PKG_VERSION}.CommitId=${GIT_COMMIT} -X ${PKG_VERSION}.BuildTime=${BUILD_TIME} -X ${PKG_VERSION}.BuildUser=${USER}"

default: pb

pb:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./proto/helloworld.proto


.PHONY: \
	all \
	test \
	release \
	clean \
