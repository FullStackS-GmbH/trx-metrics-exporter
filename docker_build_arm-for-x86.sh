#!/bin/bash

$HOME/.rd/bin/docker buildx build --platform=linux/amd64 \
 -t fullstacksgmbh/trx-metrics-exporter:$TAG \
 --no-cache \
 .



