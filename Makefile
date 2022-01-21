SHELL=/bin/bash
BINARY_NAME=alertingwebhook
VERSION=`git describe --tags --abbrev=0`
GO_VERSION=`go version|awk '{print $$3" "$$4}'`
GIT_URL=`git remote -v|grep push|awk '{print $$2}'`
GIT_BRANCH=`git rev-parse --abbrev-ref HEAD`
GIT_COMMIT=`git rev-parse HEAD`
GIT_LATEST_TAG=`git describe --tags --abbrev=0`
BUILD_TIME=`date +"%Y-%m-%d %H:%M:%S %Z"`

LDFLAGS="-X 'github.com/xmapst/alertingwebhook.Version=${VERSION}' -X 'github.com/xmapst/alertingwebhook.GoVersion=${GO_VERSION}' -X 'github.com/xmapst/alertingwebhook.GitUrl=${GIT_URL}' -X 'github.com/xmapst/alertingwebhook.GitBranch=${GIT_BRANCH}' -X 'github.com/xmapst/alertingwebhook.GitCommit=${GIT_COMMIT}' -X 'github.com/xmapst/alertingwebhook.GitLatestTag=${GIT_LATEST_TAG}' -X 'github.com/xmapst/alertingwebhook.BuildTime=${BUILD_TIME}'"

all: linux

linux:
	git pull
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags ${LDFLAGS} -o ${BINARY_NAME}_linux cmd/linux/main.go
	strip --strip-unneeded ${BINARY_NAME}_linux
	upx --lzma ${BINARY_NAME}_linux