#!/usr/bin/dumb-init /bin/bash
#
# Copyright (c) 2019 VMWare
#
# SPDX-License-Identifier: Apache-2.0
#

#Exit the script immediately if a command exits with a non-zero status
set -e

###
# Run MongoDB
###
mongod --bind_ip_all &

###
# Run Edgex-Mongo Go Application and keep the process/container alive
###
cd cmd/
./edgex-mongo --profile=docker --confdir=res
wait

