.PHONY: build docker run


docker:
	docker build \
	 --no-cache=true \
		-t my-mongo \
		. 
