
# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help


# DOCKER TASKS

current_dir = $(shell pwd)

# Build the container
build-docker: ## Build the release and develoment container.
	docker build -t mobilebff .

# Start container locally and map 'src' dir to container. .ts files are continuously build. 
dev: build-docker ## Spin up the project
	docker run -v ${current_dir}/src:/server/src --network georace_default -p 8088:8080 mobilebff npm run dev 
