.PHONY: docker 

DOCKERS=docker_x86-64
.PHONY: $(DOCKERS)

GIT_SHA=$(shell git rev-parse HEAD)

docker: $(DOCKERS)

docker_x86-64:
	 docker build \
		-f Dockerfile \
		--label "git_sha=$(GIT_SHA)" \
		-t edgexfoundry/docker-edgex-mongo:$(GIT_SHA) \
		-t edgexfoundry/docker-edgex-mongo:1.1.0 \
		.
