#! /bin/bash
set -e

mkdir /workspace/go-sdk-release

# Copy sources
cd /workspace/go-sdk-release
cp -R /builds/vapi-sdk/vsphere-automation-sdk-go ./
cd vsphere-automation-sdk-go
git remote set-url origin https://oauth2:${GO_SDK_PROJECT_ACCESS_TOKEN}@gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go.git
git remote add github https://oauth2:${GITHUB_SDK_CODE_SYNC_CI_TOKEN}@github.com/vmware/vsphere-automation-sdk-go.git

# Checkout and push
cd /workspace/go-sdk-release/vsphere-automation-sdk-go/
git checkout release
git pull origin release
git push --follow-tags github HEAD:master
