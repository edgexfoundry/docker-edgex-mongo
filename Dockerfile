###############################################################################
# Copyright 2016-2019 Dell Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
###############################################################################
# Mongo DB image for EdgeX Foundry
FROM mongo:3.4.9

RUN printf "deb http://archive.debian.org/debian/ jessie main\ndeb-src http://archive.debian.org/debian/ jessie main\ndeb http://security.debian.org jessie/updates main\ndeb-src http://security.debian.org jessie/updates main" > /etc/apt/sources.list
RUN apt-get update && apt-get install -y curl && apt-get install -y sudo
#RUN apt-get update && apt-get install -y curl && apt-get install jq


#copy initialization script for later initialization
#ADD resp-init.json /tmp/
#ADD cred.json /tmp/
ADD init.sh /docker-entrypoint-initdb.d/
#expose Mongodb's port
EXPOSE 27017
