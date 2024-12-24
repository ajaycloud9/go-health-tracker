#!/bin/bash


# Function to build Docker image, tag, save, and import
build_system() {
    local image_name=$1
    local tag=$2
    cd ../../../
    echo "Current Dir is $(pwd)"
    echo "Building Docker image for $image_name..."
    if ! ./scripts/build-docker-image.sh "$image_name" "$tag"; then
        cd -
        return 255
    fi

    echo "Tag Docker image..."
    if ! docker tag "$tag:latest" "$tag:1.0.0"; then
        cd -
        return 255
    fi

    cd -
    echo "Done for $image_name"
}

build_system "src" "go-health-backend" || false
build_system "ui" "go-health-ui" || false


BUILD_RC=$?

# LAST: Validate build completion
if [[ $BUILD_RC == 0 ]]; then
    echo -e "\nDONE: Successfully completed building all services."
else
    echo -e "\nFAIL: Failure occurred while building all services."
    exit 1
fi
