source ~/.bash_profile

cd /minitwit/scripts

docker compose -f docker-compose.yml pull
docker compose -f docker-compose.yml up -d