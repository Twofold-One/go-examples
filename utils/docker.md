# Docker cheatsheet

https://hub.docker.com/

For linux system prefix `sudo` before every cli command.

## Common practices and Postgres specific

To get all available docker images:

`$ docker images`

To pull docker img of postgres db:

`$ docker pull postgres: 14.2-alpine`

To start the container( -p <host_ports:container_ports>):

`$ docker run --name postgres14 -p 5434:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.2-alpine`

To check if port in use check which program is listening on port 8080 (list open files):

`$ sudo lsof -i :8080`

To list all containers:

`$ docker ps -a`

To remove container:

`$ docker rm <container_name>`

To get access to container db psql:

`$ docker exec -it postgres14 psql -U root`

To get logs:

`$ docker logs <container_name>`

To start existing container:

`$ docker start <container_name>`

To stop container:

`$ docker stop container_name`

## MONGODB Specific

To pull docker img of mongodb:

`$ docker pull mongo`

To start mongodb container:

`$ docker run --name mongodb -p 27017:27017 -e MONGODB_INITDB_ROOT_USERNAME=root -e MONGODB_INITDB_ROOTPASSWORD=secret -d mongo:latest`

To get access to mongodb shell:

`$ docker exec -it mongodb bash` or `$ docker exec -it mongodb mongo`
