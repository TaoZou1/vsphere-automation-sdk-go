#! /bin/bash
set -e

commitId=${1}
echo "Commit Id: $commitId"

nextVersion=${2}
echo "Next version: $nextVersion"

modulePath=${3}
echo "Module path: $modulePath"

newTag="$modulePath/$nextVersion"

mkdir /workspace/go-sdk-main

# Copy sources
cd /workspace/go-sdk-main
cp -R /builds/vapi-sdk/vsphere-automation-sdk-go ./
cd vsphere-automation-sdk-go
git remote set-url origin https://oauth2:${GO_SDK_PROJECT_ACCESS_TOKEN}@gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go.git

# Commit and push
cd /workspace/go-sdk-main/vsphere-automation-sdk-go/
git checkout main
git pull origin main
echo "$nextVersion" > /workspace/go-sdk-main/vsphere-automation-sdk-go/$modulePath/version.txt
git add $modulePath/version.txt
if [[ `git status --porcelain` ]]; then
  git commit -m "Updated version.txt"
  versionUpdateCommitId=$(git log --format="%H" -n 1)
  git push origin main

  git checkout release
  git pull origin release
  git cherry-pick $commitId
  git cherry-pick $versionUpdateCommitId
  git tag $newTag
  git push origin release
  git push origin $newTag
else
  echo "Nothing to commit."
fi

