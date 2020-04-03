#!/bin/bash

echo "$DOCKERHUB_PASSWORD" | docker login -u "$DOCKERHUB_USERNAME" --password-stdin
FULL_CHROME_VERSION=$(docker run --rm -it "$IMAGE_NAME" bash -c 'echo "$(google-chrome --version)"')
CHROME_VERSION=${FULL_CHROME_VERSION//[!0-9.]/}
CHROME_TAG="CHROME-${CHROME_VERSION}"
CIRCLECI_TAG="CIRCLECI-${CIRCLE_BUILD_NUM}"
GAUGE_VERSION="1.0.8"
GAUGE_TAG="GAUGE-$GAUGE_VERSION"
GIT_TAG="GIT-COMMIT-${CIRCLE_SHA1}"
NODE_VERSION=13.12
NODE_TAG="NODE-${NODE_VERSION}"
# Tag Gauge version as the semantic version of the image, as that is what CircleCI do
# for their base images: https://circleci.com/docs/2.0/circleci-images/#best-practices
docker tag "$IMAGE_NAME:latest" "$IMAGE_NAME:$GAUGE_VERSION"
# Also tag the Gauge version and the circle build together so that consumers can pin to an 
# idempotent image
docker tag "$IMAGE_NAME:latest" "$IMAGE_NAME:$GAUGE_VERSION-$CIRCLECI_TAG"
docker tag "$IMAGE_NAME:latest" "$IMAGE_NAME:$CHROME_TAG"
docker tag "$IMAGE_NAME:latest" "$IMAGE_NAME:$CIRCLECI_TAG"
docker tag "$IMAGE_NAME:latest" "$IMAGE_NAME:$GAUGE_TAG"
docker tag "$IMAGE_NAME:latest" "$IMAGE_NAME:$GIT_TAG"
docker tag "$IMAGE_NAME:latest" "$IMAGE_NAME:$NODE_TAG"
echo "Created the following tags:"
docker images "$IMAGE_NAME" --format="{{ .Tag }}" | tee -a  "/tmp/tags"
if [ "${CIRCLE_BRANCH}" == "master" ]; then
      docker push "$IMAGE_NAME"
fi
