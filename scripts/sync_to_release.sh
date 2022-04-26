#! /bin/bash
set -e

commitId=${1}
echo "Commit Id: $commitId"

nextVersion=${2}
echo "Next version: $nextVersion"

modulePath=${3}
echo "Module path: $modulePath"

cd /workspace/go-sdk-main/vsphere-automation-sdk-go

# Commit and push
cd /workspace/go-sdk-main/vsphere-automation-sdk-go/
echo "$nextVersion" > /workspace/go-sdk-main/vsphere-automation-sdk-go/$modulePath/version.txt
git add $modulePath/version.txt
if [[ `git status --porcelain` ]]; then
  git commit -m "Updated version.txt"
  versionUpdateCommitId=$(git log --format="%H" -n 1)
  git push origin aagrawal3/main/automate-sementic-versioning

  git checkout aagrawal3/main/test-sync-to-release
  git pull origin aagrawal3/main/test-sync-to-release
  git cherry-pick $commitId
  git cherry-pick $versionUpdateCommitId
  git push origin aagrawal3/main/test-sync-to-release
else
  echo "Nothing to commit."
fi

