echo "Clear compose"

docker compose stop
docker compose rm -f

echo "Clear images"
docker image rm testtask-go-backend
docker image rm testtask-go-msg-handler
docker image rm testtask-db
docker image rm testtask-kafka
docker image rm testtask-zookeeper

