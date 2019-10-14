DOCKER = /usr/bin/docker
BUILD_ARG = $(if $(filter  $(NOCACHE), 1),--no-cache)

production: production_destroy_disk_volumes code_image development_code
staging: staging_destroy_disk_volumes code_image development_code
development: develop_destroy_disk_volumes code_image development_code
clean: volumes_down destroy_disk_volumes
develop_stop: develop_shutdown develop_volumes_down
production_stop: production_shutdown production_volumes_down
staging_stop: staging_shutdown staging_volumes_down
production_up: production_start
staging_up: staging_start
develop_up: develop_start


production_start:
	docker-compose -f deployments/docker-compose.production.yml up -d
staging_start:
	docker-compose -f deployments/docker-compose.staging.yml up -d
develop_start:
	docker-compose -f deployments/docker-compose.yml up -d
develop_shutdown:
	docker-compose stop
develop_volumes_down:
	docker-compose deployments/docker-compose.yml down --volumes
staging_shutdown:
	docker-compose -f deployments/docker-compose.staging.yml stop
staging_volumes_down:
	docker-compose -f deployments/docker-compose.staging.yml down --volumes
production_shutdown:
	docker-compose -f deployments/docker-compose.production.yml stop
production_volumes_down:
	docker-compose -f deployments/docker-compose.production.yml down --volumes
production_destroy_disk_volumes:
	$(DOCKER) volume rm -f j-vazquez.com
staging_destroy_disk_volumes:
	$(DOCKER) volume rm -f jvazquez.xyz
develop_destroy_disk_volumes:
	$(DOCKER) volume rm -f develop.jvazquez
production_disk_volumes:
	$(DOCKER) volume create j-vazquez.com
staging_disk_volumes:
	$(DOCKER) volume create jvazquez.xyz
staging_disk_volumes:
	$(DOCKER) volume create develop.jvazquez
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
