IMAGE_TAG=$(./version.sh)

images=$(cat docker-compose.yaml | grep 'image: moveyourfeet' | cut -d':' -f 2 | tr -d '"')
for image in $images
do
  docker tag "${image}":latest "${image}":"${IMAGE_TAG}"
  docker push "${image}":"${IMAGE_TAG}"
done