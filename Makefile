.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose kill && docker-compose rm -f

.PHONY: dev
dev:
	docker-compose up -d --build

.PHONY: build
build:
	docker build . -t katsew/kodama:latest
