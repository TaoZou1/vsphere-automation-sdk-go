#! /bin/bash
set -e
set -x
export GOPROXY="https://packages.vcfd.broadcom.net/artifactory/proxy-golang-remote"
export GOSUMDB="sum.golang.org https://packages.vcfd.broadcom.net/artifactory/go-gosumdb-remote"

# get the module path
echo $CI_COMMIT_TAG
modulePath=$(echo $CI_COMMIT_TAG | rev | cut -d "/" -f 2- | rev)
echo $modulePath

echo "Printing current working directory..."
pwd

# cleaning directory(if already exists) before cloning github repo
if [ -d "$DESTINATION_REPO" ] 
then 
    rm -rf $DESTINATION_REPO 
fi

# hardcoding master branch to avoid accidental tags in other branches
git clone -b master --single-branch https://oauth2:${CI_ACCESS_TOKEN}@${DESTINATION_REPO}.git $DESTINATION_REPO
cd $DESTINATION_REPO
echo "Listing files from destination repo root dir: "
ls -la

moduleRepoPath="${DESTINATION_REPO}/$modulePath/"

# clear the module before copying 
# do this only if directory exists
if [ -d "$modulePath" ] 
then
    cd $modulePath
    rm -rf $(go list ./... | grep "$moduleRepoPath" | sed "s#$moduleRepoPath##")

    # deleting empty directories
    find . -empty -type d -delete
    # removing root level go files(if any)
    rm -f go.mod go.sum *.go

    # listing the modulePath before copying 
    echo "Printing files in $modulePath after cleanup: "
    ls -la 
fi

# coming out of the github repo and going inside the gitlab repo scripts directory to run the gradle task
cd $CI_PROJECT_DIR/scripts/gradle
./gradlew copyModule -PmodulePath=$modulePath -PdestRepoGithub=$DESTINATION_REPO

echo "Printing files in $modulePath after copy: "
ls -la $CI_PROJECT_DIR/$DESTINATION_REPO/$modulePath
