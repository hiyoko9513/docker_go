up-build:
	docker-compose up -d --build
up:
	docker-compose up -d
build:
	docker-compose build --no-cache --force-rm
stop:
	docker-compose stop
down:
	docker-compose down
restart:
	@make down
	@make up
main-go:
	docker-compose exec go go run main.go
realize-start:
	docker-compose exec go realize start
