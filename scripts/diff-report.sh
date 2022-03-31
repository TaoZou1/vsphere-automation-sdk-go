
#! /bin/bash
set -e

modulePath=${1:-}
echo "Module path: $modulePath"

mkdir /workspace/go-sdk-main
mkdir /workspace/go-sdk-tag

cd /workspace/go-sdk-main

# Clone repo
git clone https://oauth2:${SDK_CLONE_ACCESS_TOKEN}@gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go.git

echo "Copy repo in tag dir..."
cp -R /workspace/go-sdk-main/vsphere-automation-sdk-go /workspace/go-sdk-tag/

cd /workspace/go-sdk-main/vsphere-automation-sdk-go/
git checkout main
git pull origin main

cd /workspace/go-sdk-tag/vsphere-automation-sdk-go/
# Get new tags from remote
git fetch --tags

# Get latest tag name
latestTag=$(git describe --match "$modulePath*" --tags `git rev-list --tags --max-count=1`)
echo "Last tagged version for $modulePath was $latestTag"

# Identify the service
# modulePath=$(echo $latestTag | rev | cut -d "/" -f 2- | rev)
mkdir -p /workspace/results/$modulePath
ls -la /workspace/results/$modulePath

# Checkout latest tag
git checkout $latestTag

# run diff check
cd /workspace/diff-checker
echo "Old mod files"
ls -la /workspace/go-sdk-tag/vsphere-automation-sdk-go/$modulePath
echo "New mod files"
ls -la /workspace/go-sdk-main/vsphere-automation-sdk-go/$modulePath

echo "Generating diff..."
go run cmd/module-diff-check.go generate-report --o /workspace/go-sdk-tag/vsphere-automation-sdk-go/$modulePath --n /workspace/go-sdk-main/vsphere-automation-sdk-go/$modulePath --result-dir /workspace/results/$modulePath --lang go
cat /workspace/results/$modulePath/go-mod-final-report.json
