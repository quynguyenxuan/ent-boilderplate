cd "$(dirname "$0")"

docker stack deploy -c docker-compose.yml whm --with-registry-auth