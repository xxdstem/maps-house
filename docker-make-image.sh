#! /bin/bash
docker buildx build \
                --platform linux/arm64 \
                --output "type=image,push=true" \
                --tag xxdstem/maps-house .