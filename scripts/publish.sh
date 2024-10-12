#! /bin/bash
set -e
set -x 

modulePath=$(echo $CI_COMMIT_TAG | rev | cut -d "/" -f 2- | rev)

cd $CI_PROJECT_DIR/$DESTINATION_REPO/$modulePath

git log -5 --no-walk --tags --pretty="%h %d %s" --decorate=full
REPO="https://oauth2:${CI_ACCESS_TOKEN}@${DESTINATION_REPO}.git"
git push --follow-tags $REPO HEAD:master