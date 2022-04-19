#! /bin/bash
set -e

# Clone repo
cd /workspace/go-sdk-main
cp -R /builds/vapi-sdk/vsphere-automation-sdk-go ./
cd vsphere-automation-sdk-go
git remote set-url origin https://oauth2:${SDK_CLONE_ACCESS_TOKEN}@gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go.git

# Checkout respective branch codes
cd /workspace/go-sdk-main/vsphere-automation-sdk-go/
git checkout aagrawal3/main/automate-sementic-versioning
git pull origin aagrawal3/main/automate-sementic-versioning

# git checkout release
# git pull origin release
git checkout -b aagrawal3/main/test-sync-to-release
git merge aagrawal3/main/automate-sementic-versioning
git push origin aagrawal3/main/test-sync-to-release

