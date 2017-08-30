#!/bin/sh

set -e

mongod --smallfiles &

sleep 10

mongo /edgex/mongo/config/init_mongo.js

wait

