#! /bin/bash
tag=$1
git pull
git checkout $tag
docker buildx build . \
--platform linux/amd64,linux/arm64,linux/arm/v7 \
-t dezhishen/onebot-plus:$tag --push