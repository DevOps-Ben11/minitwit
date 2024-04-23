source ~/.bash_profile

cd /minitwit/scripts

docker compose -f docker-compose.yml pull
docker stack deploy ministack --compose-file docker-compose.prod.yml
