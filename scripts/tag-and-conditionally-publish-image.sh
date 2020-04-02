#!/bin/bash

echo "$DOCKERHUB_PASSWORD" | docker login -u "$DOCKERHUB_USERNAME" --password-stdin
CIRCLECI_TAG="CIRCLECI-${CIRCLE_BUILD_NUM}"
GAUGE_VERSION="1.0.8"
GAUGE_TAG="GAUGE-$GAUGE_VERSION"
GIT_TAG="GIT-COMMIT-${CIRCLE_SHA1}"
NODE_VERSION=13.12
NODE_TAG="NODE-${NODE_VERSION}"
echo "Using the Gauge version as the semantic version of the image, as that is what CircleCI do
      for their base images: https://circleci.com/docs/2.0/circleci-images/#best-practices"
docker tag "$IMAGE_NAME:latest" "$IMAGE_NAME:$GAUGE_VERSION"
docker tag "$IMAGE_NAME:latest" "$IMAGE_NAME:$CIRCLECI_TAG"
docker tag "$IMAGE_NAME:latest" "$IMAGE_NAME:$GAUGE_TAG"
docker tag "$IMAGE_NAME:latest" "$IMAGE_NAME:$GIT_TAG"
docker tag "$IMAGE_NAME:latest" "$IMAGE_NAME:$NODE_TAG"
if [ "${CIRCLE_BRANCH}" == "master" ]; then
      docker push "$IMAGE_NAME"
fi
