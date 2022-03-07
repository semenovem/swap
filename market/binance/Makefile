include project.properties

IMG := "$(DOCKER_IMAGE_NAME):$(VERSION)"
IMG_DEV := "$(DOCKER_IMAGE_NAME)-dev:$(VERSION)"

# -----------------
# commands
# -----------------

.PHONY: help
help: Makefile
	@echo "Choose a command in:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'


## docker
.PHONY: docker
docker:
	@go mod vendor
	@docker build \
		--build-arg BUILD_VERSION="$(VERSION)" \
		--build-arg BUILD_HASH_COMMIT="$(shell git rev-parse --verify HEAD)" \
		--build-arg BUILD_BRANCH="$(shell git rev-parse --abbrev-ref HEAD)" \
		--build-arg BUILD_TIME="$(shell date)" \
 		-f ./Dockerfile \
 		-t "$(IMG)" .
	@rm -rf vendor


## registry
.PHONY: registry
registry:
	@make docker
	@docker image tag $(IMG) $(REGISTRY)/$(IMG)
	@docker push $(REGISTRY)/$(IMG)


## dev - контейнер для разработки. Запуск приложения внутри контейнера: `dev/run.sh`
.PHONY: dev
dev:
	@#bash dev/check-img.sh "$(IMG_SOFTHSM)" "$(IMG_SOFTHSM_DEV)" `pwd`/..
	@docker run --rm -it \
		--network=$(NETWORK) \
		--name binance \
		--hostname server \
		-w /app/api/binance \
		-v `pwd`/../..:/app:rw \
		-p 7000:2022 \
		"$(IMG)" bash dev/run.sh
