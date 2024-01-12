help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

OUTPUT_REPO=nubuscu.github.io

run-backend:  ## use docker-compose to run the api/db pods
	docker-compose up --build --remove-orphans

run-frontend-dev:  ## run the frontend in hot-relaod/dev mode
	cd vis && npm run dev

build:  ## build the frontend
	docker-compose up --build --remove-orphans --detach
	cd vis && BASE_PATH=$(OUTPUT_REPO) npm run build

preview: ## preview the latest build of the frontend
	cd vis && npm run preview

copy-to-deploy: ## use with message=commit message
	rm -rf $(OUTPUT_REPO)/build/
	cp -r vis/build/ $(OUTPUT_REPO)/build/