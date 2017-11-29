.PHONY: image
image:
	./bin/docker-image.sh

.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose kill && docker-compose rm -f

.PHONY: dev
dev:
	docker-compose up -d --build
