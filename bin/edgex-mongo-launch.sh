#!/usr/bin/dumb-init /bin/bash
#
# Copyright (c) 2019 VMWare
#
# SPDX-License-Identifier: Apache-2.0
#

#Exit the script immediately if a command exits with a non-zero status
set -e

#sentinel file used to determine if the database was already initialized
: ${DB_INITIALIZATION_FLAG:=/data/db/.edgex-mongo-database-setup-done}

if [ ! -f "${DB_INITIALIZATION_FLAG}" ]; then
    # Run Mongo DB bind to localhost.
    mongod &

    echo "Run database initialization process"
    cd cmd/
    ./edgex-mongo --profile=docker --confdir=res

    #Shutdown mongo and later to be started up with enabled authentication
    mongod --shutdown

    echo "Signaling mongo database initialization process is completed"
    mkdir -p `dirname "${DB_INITIALIZATION_FLAG}"`
    touch "${DB_INITIALIZATION_FLAG}"
else
   echo "Database already has been initialized"
fi

# Start Mongo DB with enabled authentication and bind_ip_all
mongod --auth --bind_ip_all &
wait

