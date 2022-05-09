<!-- to pull docker img of postgres db -->

$ docker pull postgres: 14.2-alpine

% to start the container
% -p <host_ports:container_ports>
$ docker run --name postgres14 -p 5434:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.2-alpine

<!-- to check if port in use -->
<!-- check which program is listening on port 8080
list open files
$sudo lsof -i :8080 -->

<!-- to list all containers
$ docker ps -a -->

<!-- to remove container
$ docker rm <container_name> -->

<!-- to get access to container
$ docker exec -it postgres14 psql -U root -->

<!-- to get logs
$ docker logs <container_name> -->
