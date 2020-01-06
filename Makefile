###############################################################################
# Copyright 2019 VMWare.
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

.PHONY: docker build clean

DOCKERS=edgex-mongo
.PHONY: $(DOCKERS)

MICROSERVICES=cmd/edgex-mongo
.PHONY: $(MICROSERVICES)

GO=CGO_ENABLED=0 GO111MODULE=on go

VERSION=$(shell cat ./VERSION 2>/dev/null || echo 0.0.0)
DOCKER_TAG=$(VERSION)

GOFLAGS=-ldflags "-X github.com/edgexfoundry/docker-edgex-mongo.Version=$(VERSION)"

GIT_SHA=$(shell git rev-parse HEAD)

build: $(MICROSERVICES)

cmd/edgex-mongo:
	$(GO) build $(GOFLAGS) -o $@ ./cmd

docker: $(DOCKERS)

clean:
	rm -f $(MICROSERVICES)

test:
	GO111MODULE=on go test -coverprofile=coverage.out ./...
	GO111MODULE=on go vet ./...
	gofmt -l .
	./bin/test-go-mod-tidy.sh
	./bin/test-attribution-txt.sh

run:
	./bin/edgex-mongo-launch.sh

edgex-mongo:
	 docker build \
		-f cmd/Dockerfile \
		--label "git_sha=$(GIT_SHA)" \
		-t edgexfoundry/docker-edgex-mongo:$(GIT_SHA) \
		-t edgexfoundry/docker-edgex-mongo:$(DOCKER_TAG) \
		.
