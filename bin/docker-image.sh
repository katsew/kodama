#! /bin/bash

# Change directory to project root.
cd `dirname $0`
cd ../

# Initialize configurations.

## Remove comment out and set repo name.
# REPO_NAME=[repo name]

if [ -z $REPO_NAME ]; then
    echo "No repo name set!"
    exit 1
fi
VERSION=`git tag  -l --sort "-v:refname" | head -n 1`
IMAGE_NAME=kodama
mkdir -p ./tmp/bin

if [ $1 = aws ]; then

# Login to docker repo for aws
aws ecr get-login --no-include-email --region ap-northeast-1 > ./tmp/bin/login.sh
chmod 744 ./tmp/bin/login.sh
./tmp/bin/login.sh

fi

# Checkout to latest version
echo "Checking out branch to tag: $VERSION"
git checkout ${VERSION}

# Build image and push to repo
docker build -t ${IMAGE_NAME}:${VERSION} .
docker tag ${IMAGE_NAME}:${VERSION} ${REPO_NAME}/${IMAGE_NAME}:${VERSION}
docker push ${REPO_NAME}/${IMAGE_NAME}:${VERSION}
docker push ${REPO_NAME}/${IMAGE_NAME}:latest

# Cleanup
rm -r ./tmp
git checkout master
