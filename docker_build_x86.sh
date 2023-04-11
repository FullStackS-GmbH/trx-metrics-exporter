#!/bin/bash

docker build \
 -t registry.fullstacks.eu/fullstacks/trx-metrics-exporter:$TAG \
 --no-cache \
 .
