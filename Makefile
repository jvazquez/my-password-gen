DOCKER = /usr/bin/docker
BUILD_ARG = $(if $(filter  $(NOCACHE), 1),--no-cache)
DEFAULT_ENVIRONMENT = development
ENVIRONMENT_TO_BUILD := $(if $(ENVIRONMENT),$(ENVIRONMENT),$(DEFAULT_ENVIRONMENT))
DOCKERFILE_PATH = $(if $(filter  $(ENVIRONMENT_TO_BUILD), production),images/python/production/Dockerfile,images/python/development/Dockerfile)
VOLUME = my-password-gen-go

development: destroy_disk_volumes code_image development_code
clean: volumes_down destroy_disk_volumes
stop: shutdown volumes_down

shutdown:
	docker-compose stop
volumes_down:
	docker-compose down --volumes
destroy_disk_volumes:
	$(DOCKER) volume rm -f $(VOLUME)
disk_volumes:
	$(DOCKER) volume create $(VOLUME)
development_code:
	$(DOCKER) build $(BUILD_ARG) -f build/go/Dockerfile -t local-my-password-gen .
	$(DOCKER) run  --rm -v $(VOLUME):/my-password-gen --name data-container local-my-password-gen bash -c 'cd /my-password-gen/cmd;\
	 go build xkcd.go'
code_image:
	$(DOCKER) volume create $(VOLUME)
	$(DOCKER) stop data-container || true && docker rm data-container || true
	$(DOCKER) run -v $(VOLUME):/my-password-gen --name data-container busybox true
	$(DOCKER) cp . data-container:/my-password-gen
	$(DOCKER) stop data-container
	$(DOCKER) rm data-container
