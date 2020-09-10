#!/bin/sh

set -e
export GO111MODULE=on
GOOS=linux go build -o reviews
docker build -t devnryn/reviews:v3 .
docker push devnryn/reviews:v3
rm reviews
