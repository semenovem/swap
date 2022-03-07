include ../project.properties

#IMG := "$(DB_IMG):$(DB_VER)"
#IMG_DEV := "$(DB_IMG)-dev:$(DB_VER)-dev"

IMG := "golang:1.17"
IMG_DEV := "$(DB_IMG)-dev:$(DB_VER)-dev"

# -----------------
# commands
# -----------------

.PHONY: help
help: Makefile
	@echo "Choose a command in:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'


## dev - контейнер для разработки. Запуск приложения внутри контейнера: `dev/run.sh`
.PHONY: dev
dev:
	@$(eval NETWORK=$(shell bash docker-network-selection.sh))
	@echo "\n=> NETWORK=$(NETWORK)"
	@docker run --rm -it \
		--network=$(NETWORK) \
		--name app \
		--hostname a \
		-w /repo/app \
		-v $(PWD)/../:/repo:rw \
		"$(IMG)" bash ../dev/run.sh


## rm-dev - удалит контейнер для разработки
.PHONY: rm-dev
rm-dev:
	@$(eval id=$(shell docker images --filter=reference="$(IMG_DEV)" -q))
	@docker image rm $(id)


## curl - контейнер, присоединенный к сети, для выполнения curl запросов
.PHONY: curl
curl:
	@$(eval NETWORK=$(shell bash docker-network-selection.sh))
	@echo "\n=> NETWORK=$(NETWORK)"
	@docker run -it --rm --network $(NETWORK) curlimages/curl:7.78.0 sh -c \
     "echo '| Для выхода из контейнера ctrl+D' && export PS1='> ' && sh"
