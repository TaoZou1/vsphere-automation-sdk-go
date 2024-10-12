#! /bin/bash
set -e
set -x

echo "Building for GIT TAG: $CI_COMMIT_TAG"
modulePath=$(echo $CI_COMMIT_TAG | rev | cut -d "/" -f 2- | rev)

ls -la ${CI_PROJECT_DIR}

echo $modulePath
cd  ${CI_PROJECT_DIR}/$DESTINATION_REPO/$modulePath
echo "List files in destination module after copy: "
ls -la

echo "Overriding Go environment variable."
export GO111MODULE=on
export GONOPROXY="github.com/vmware"
export GONOSUMDB="github.com/vmware"
export GOPROXY="https://packages.vcfd.broadcom.net/artifactory/proxy-golang-remote"
export GOSUMDB="sum.golang.org https://packages.vcfd.broadcom.net/artifactory/go-gosumdb-remote"

echo "Clean go cache"
go clean -modcache -cache
echo "Remove old go.sum file"
rm -f go.sum
echo "Running go mod tidy"
go mod tidy
echo "Building Go Module"
go build ./...
