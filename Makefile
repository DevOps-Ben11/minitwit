connect-stage:
	ssh root@46.101.214.108

connect-prod:
	ssh root@138.68.126.8

local: 
	export DOCKER_USERNAME=hrallil
	export PSQL_CON_STR="postgresql://postgres:mysecretpassword@host.docker.internal:5431/postgres"
	docker run --name my_postgres_db -e POSTGRES_PASSWORD=mysecretpassword -d -p 5431:5432 postgres
	docker compose -f scripts/docker-compose.yml up

