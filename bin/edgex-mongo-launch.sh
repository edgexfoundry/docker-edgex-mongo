#!/usr/bin/dumb-init /bin/bash
#
# Copyright (c) 2019 VMWare
#
# SPDX-License-Identifier: Apache-2.0
#

#Exit the script immediately if a command exits with a non-zero status
set -e

###
# Run MongoDB bind to localhost.
###
mongod &

###
# Run Edgex-Mongo Go Application and initiate the database
###
cd cmd/
./edgex-mongo --profile=docker --confdir=res


###
# Restart Edgex-Mongo with enabled authentication and bind_ip_all
###
mongod --shutdown
mongod --auth --bind_ip_all &
wait

