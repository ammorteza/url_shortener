compose:
	sudo docker-compose up --build

help:
	sudo docker exec urlshortener_go ./cmd/main --help

router_start:
	sudo docker exec urlshortener_go ./cmd/main router:start

db_migrate:
	sudo docker exec urlshortener_go ./cmd/main db:migrate

db_reset:
	sudo docker exec urlshortener_go ./cmd/main db:reset

api_test:
	sudo docker exec urlshortener_go go test ./test/