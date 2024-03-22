source ~/.bash_profile

cd /minitwit/scripts

docker compose -f docker-compose.yml pull

# Stop all old docker containers from running to run the new one
docker stop $(docker ps -a -q)
docker compose -f docker-compose.yml up -d