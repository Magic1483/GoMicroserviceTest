#!/bin/bash
exec ./docker_clear.sh &

echo "Build msg handler"
cd ./msg_handler
rm ./msgHandler
go build
echo "MSG HANDLER BUILDED"
cd ..

echo "Build server"
rm ./microService
go build
echo "SERVER BUILDED"

echo "docker compose up"
docker compose up


echo |----------------|
echo |    COMPLETE    |
echo |----------------|

