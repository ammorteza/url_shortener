run:
	sudo docker-compose up --build

db_migrate:
	sudo docker exec urlshortener_go ./Cli db:migrate

db_reset:
	sudo docker exec urlshortener_go ./Cli db:reset