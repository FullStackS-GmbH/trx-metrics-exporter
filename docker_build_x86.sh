#!/bin/bash

docker build \
 -t fullstacksgmbh/trx-metrics-exporter:$TAG \
 --no-cache \
 .
