
#! /bin/bash
set -e

modulePath=${1:-}
echo "Module path: $modulePath"

mkdir /workspace/go-sdk-main
mkdir /workspace/go-sdk-tag
mkdir -p /workspace/results/$modulePath

# Copy sources
cd /workspace/go-sdk-main
cp -R /builds/vapi-sdk/vsphere-automation-sdk-go ./
cd vsphere-automation-sdk-go
git remote set-url origin https://oauth2:${GO_SDK_PROJECT_ACCESS_TOKEN}@gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go.git
echo "Copy repo in tag dir..."
cp -R /workspace/go-sdk-main/vsphere-automation-sdk-go /workspace/go-sdk-tag/

# Checkout respective branch codes
cd /workspace/go-sdk-main/vsphere-automation-sdk-go/
git checkout aagrawal3/main/automate-sementic-versioning
git pull origin aagrawal3/main/automate-sementic-versioning
cd /workspace/go-sdk-tag/vsphere-automation-sdk-go/
git fetch --tags
latestTag=$(git describe --match "$modulePath*" --tags `git rev-list --tags --max-count=1`)
echo "Last tagged version for $modulePath was $latestTag"
git checkout $latestTag

# run diff check
cd /workspace/diff-checker
echo "Generating diff..."
go run cmd/module-diff-check.go generate-report --o /workspace/go-sdk-tag/vsphere-automation-sdk-go/$modulePath --n /workspace/go-sdk-main/vsphere-automation-sdk-go/$modulePath --result-dir /workspace/results/$modulePath --lang go

# find release type
cat /workspace/results/$modulePath/go-mod-final-report.json
# TODO move to docker image
apt-get install -y jq
RELEASE_TYPE=$(jq '.ReleaseType' /workspace/results/$modulePath/go-mod-final-report.json)
echo "Detected release type: $RELEASE_TYPE"

# Update version file in main branch...
cd /workspace/go-sdk-main/vsphere-automation-sdk-go/
lastReleaseVersion="${latestTag##*/}"
versionArray=($(echo $lastReleaseVersion | tr '.' "\n"))
nextRelease=$lastReleaseVersion

if [[ $RELEASE_TYPE  ==  '"MAJOR"' ]]
then
  # remove 'v' from beginning
  majorVersion=${versionArray[0]:1}
  majorVersion=$(( majorVersion + 1 ))
  nextRelease="v$majorVersion.${versionArray[1]}.${versionArray[2]}"
  echo "Next release version: $nextRelease"
elif [[ $RELEASE_TYPE  ==  '"MINOR"' ]]
then
  minorVersion=${versionArray[1]}
  minorVersion=$(( minorVersion + 1 ))
  nextRelease="${versionArray[0]}.$minorVersion.${versionArray[2]}"
  echo "Next release version: $nextRelease"
elif [[ $RELEASE_TYPE  ==  '"PATCH"' ]]
then
  patchVersion=${versionArray[2]}
  patchVersion=$(( patchVersion + 1 ))
  nextRelease="${versionArray[0]}.${versionArray[1]}.$patchVersion"
  echo "Next release version: $nextRelease"
else
  echo "no change detected..."
fi

echo "$nextRelease" > /workspace/go-sdk-main/vsphere-automation-sdk-go/$modulePath/version.txt

# Commit and push
cd /workspace/go-sdk-main/vsphere-automation-sdk-go/
git add $modulePath/version.txt
if [[ `git status --porcelain` ]]; then
  git commit -m "Updated version.txt"
  git push origin aagrawal3/main/automate-sementic-versioning
else
  echo "Nothing to commit."
fi
