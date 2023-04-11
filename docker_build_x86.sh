#!/bin/bash

docker build \
 -t fullstacksgmbh/snmp2prom:$TAG \
 --no-cache \
 .
