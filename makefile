
# grep the version from the mix file
VERSION=$(shell ./version.sh)


# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help


# DOCKER TASKS

# Build the container
build: ## Build the release and develoment container.
	docker-compose build

# Build and run the container
up: ## Spin up the project
	docker-compose up --build -d

stop: ## Stop running containers
	docker-compose stop

down: stop ## Stop and remove running containers
	docker-compose down

logs: ## Show logs for running containers
	docker-compose logs -f

clean-build: ## Clean the generated/compiles files
	docker-compose build --no-cache

docs: up ## Open documentation page in browser
	$(shell open docs.localtest.me/ || sensible-browser docs.localtest.me/)

# Docker release - build, tag and push the container
# release: build publish ## Make a release by building and publishing the `{version}` ans `latest` tagged containers to ECR

# Docker publish
# publish: repo-login publish-latest publish-version ## publish the `{version}` ans `latest` tagged containers to ECR

# publish-latest: tag-latest ## publish the `latest` taged container to ECR
# 	@echo 'publish latest to $(DOCKER_REPO)'
# 	docker push $(DOCKER_REPO)/$(APP_NAME):latest

# publish-version: tag-version ## publish the `{version}` taged container to ECR
# 	@echo 'publish $(VERSION) to $(DOCKER_REPO)'
# 	docker push $(DOCKER_REPO)/$(APP_NAME):$(VERSION)

# Docker tagging
# tag: tag-latest tag-version ## Generate container tags for the `{version}` ans `latest` tags

# tag-latest: ## Generate container `{version}` tag
# 	@echo 'create tag latest'
# 	docker tag $(APP_NAME) $(DOCKER_REPO)/$(APP_NAME):latest

# tag-version: ## Generate container `latest` tag
# 	@echo 'create tag $(VERSION)'
# 	docker tag $(APP_NAME) $(DOCKER_REPO)/$(APP_NAME):$(VERSION)


# HELPERS

# generate script to login to aws docker repo
# CMD_REPOLOGIN := "aws ecr"
# ifdef AWS_CLI_PROFILE
# CMD_REPOLOGIN += "--profile $(AWS_CLI_PROFILE)"
# endif
# ifdef AWS_CLI_REGION
# CMD_REPOLOGIN += "--region $(AWS_CLI_REGION)"
# endif
# CMD_REPOLOGIN += "get-login --no-include-email"

# repo-login: ## Auto login to AWS-ECR unsing aws-cli
# 	@eval $(CMD_REPOLOGIN)

version: ## output to version
	@echo $(VERSION)

