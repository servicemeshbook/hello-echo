VERSION=1.0
REPO=vikramkhatri
IMAGE=gorest
TAG=$(REPO)/$(IMAGE):$(VERSION)

.PHONY: help
help: ## - Show help message
	@printf "\033[32m\xE2\x9c\x93 usage: make [target]\n\n\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build:	## - Build the docker image
	@printf "\033[32m\xE2\x9c\x93 Build the docker image \n\033[0m"
	@export DOCKER_CONTENT_TRUST=1 && sudo docker build --no-cache -f Dockerfile -t $(TAG) .

.PHONY: run
run:	## - Run the golang docker image 
	@printf "\033[32m\xE2\x9c\x93 Run the docker image\n\033[0m"
	@sudo sudo docker run -d --name=$(IMAGE) -p 8080:8080 $(TAG)

.PHONY: push
push:	## - Push docker image container registry
	@sudo docker login
	@sudo docker push $(TAG)

.PHONY: clean
clean:	## - Remove docker image and go binary
	@sudo docker stop $(IMAGE)
	@sudo docker rm $(IMAGE)
	@sudo docker rmi $(TAG)
	@rm -f main
