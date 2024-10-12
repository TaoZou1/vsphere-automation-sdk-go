#! /bin/bash
set -e
set -x 

modulePath=$(echo $CI_COMMIT_TAG | rev | cut -d "/" -f 2- | rev)

cd $CI_PROJECT_DIR/$DESTINATION_REPO/$modulePath

git status
COMMIT_MESSAGE="Updated client bindings for $CI_COMMIT_TAG. Signed-off-by: sdk team <dx-sdk-team@vmware.com>"

echo "Adding changed files to git staging."
git add .

echo "Check if there are any changes to commit."
set +e 
git diff --exit-code --cached > /dev/null
diffResult=$?
set -e

if [ "$diffResult" = "0" ]
then
    echo "Nothing to commit." 
    exit 0
fi

echo "Checking if Tag $CI_COMMIT_TAG exists."
if [ $(git tag -l "$CI_COMMIT_TAG") ]; then
    echo "Tag $CI_COMMIT_TAG exists."
    exit 1
fi

echo "Commiting changes."
git commit -m "$COMMIT_MESSAGE"
echo "Updating client bindings for $CI_COMMIT_TAG"
git tag -a -m "$CI_COMMIT_TAG" $CI_COMMIT_TAG
git log -5 --no-walk --tags --pretty="%h %d %s" --decorate=full
