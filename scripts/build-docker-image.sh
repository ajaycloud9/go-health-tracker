#!/usr/bin/env bash

set -ex

DOCKER_TAG="go-health-app:1.0.0"
CTXT="DockerContext-${DOCKER_TAG}"


cleanup() {
	# Golang making rrr directories...
	find . -type d -print | xargs chmod u+w
	cd ..
	rm -rf "$CTXT"
}

if [ -e "$CTXT" ]; then
	rm -rf "$CTXT"
fi

mkdir "$CTXT" # make clean
cd "$CTXT"

# make sure we're in the right dir first before registering the cleanup handler
trap "cleanup" EXIT

for i in src common; do
	if [ ! -e "${i}" ]; then
    pwd
    cp -al ../${i} .
	fi
done


echo "BUILDING DOCKER"
pwd
docker build -f "src/Dockerfile" -t "${DOCKER_TAG}" .
