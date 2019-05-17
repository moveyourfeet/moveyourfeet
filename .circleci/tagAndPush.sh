#!/bin/bash
IMAGE_TAG="$(git describe --tags)"

if [ "${CIRCLE_BRANCH}" == "master" ]; then
  echo "$DOCKERHUB_PASS" | docker login -u "$DOCKERHUB_USERNAME" --password-stdin > /dev/null 2>&1
    
  images=$(cat docker-compose.yaml | grep "image: $CIRCLE_PROJECT_USERNAME" | cut -d':' -f 2 | tr -d '"')

  for image in $images
  do
    docker tag "${image}":latest "${image}":"${IMAGE_TAG}"
    echo "Pushing : ${image}":"${IMAGE_TAG}"
    docker push "${image}":"${IMAGE_TAG}"
  done
fi
