# version: 1.0
source ~/.bash_profile

cd /minitwit/scripts
docker stack deploy ministack --compose-file docker-compose.prod.yml
