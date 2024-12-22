#!/usr/bin/env bash

set -ex

if [ "$#" -ne 2 ]; then
	echo "Please supply the microservice name and docker tag."
	exit 1
fi

SERVICE="$1"
DOCKER_TAG="$2"
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

for i in "$SERVICE" common; do
	if [ ! -e "${i}" ]; then
    pwd
    cp -al ../${i} .
	fi
done


echo "BUILDING DOCKER"
pwd
docker build -f "${SERVICE}/Dockerfile" -t "${DOCKER_TAG}" .
