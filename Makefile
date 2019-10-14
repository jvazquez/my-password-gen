DOCKER = /usr/bin/docker
BUILD_ARG = $(if $(filter  $(NOCACHE), 1),--no-cache)
VOLUME = my-password-gen-go

development: destroy_disk_volumes code_image development_code
clean: volumes_down destroy_disk_volumes
develop_stop: develop_shutdown develop_volumes_down
production_stop: production_shutdown production_volumes_down
staging_stop: staging_shutdown staging_volumes_down
production_up: production_start
staging_up: staging_start
develop_up: develop_start
production_start:
	docker-compose -f docker-compose.production.yml up -d
staging_start:
	docker-compose -f docker-compose.staging.yml up -d
develop_start:
	docker-compose -f docker-compose.yml up -d
develop_shutdown:
	docker-compose stop
develop_volumes_down:
	docker-compose down --volumes
staging_shutdown:
	docker-compose -f docker-compose.staging.yml stop
staging_volumes_down:
	docker-compose -f docker-compose.staging.yml down --volumes
production_shutdown:
	docker-compose -f docker-compose.production.yml stop
production_volumes_down:
	docker-compose -f docker-compose.production.yml down --volumes
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
