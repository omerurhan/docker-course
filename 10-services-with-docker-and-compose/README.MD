# Services with Docker and Compose

## A Service

````
# get docker swarm status
docker info | grep -i swarm
docker node ls

# enables the service abstraction
docker swarm init

# create service
docker service create \
 --publish 8080:80 \
 --name hello-world \
 omerurhan/dockercourse:ch11_service_hw_v1

# list services and containers
docker service ls
docker ps
 
# delete container and observe
docker container rm -f <container-id>
docker ps

# details of service
docker service inspect hello-world

````
### Automated resurrection and replication

````
# replicate hello-world service
docker service scale hello-world=3
docker service ps hello-world

# scale-down and observe
docker service scale hello-world=2
docker service scale hello-world=3
````


### Automated rollout
````
# update image
docker service update \
 --image omerurhan/dockercourse:ch11_service_hw_v2 \
 --update-order stop-first \
 --update-parallelism 1 \
 --update-delay 30s \
 hello-world 
````
### Service health and rollback
````
docker service ps hello-world

docker service update \
--image omerurhan/dockercourse:ch11_service_hw_start_failure \
hello-world

# rollback to previous desired state
docker service update \
--rollback \
hello-world

# update-failure-action
docker service update \
--update-failure-action rollback \
--update-max-failure-ratio 0.6 \
--image omerurhan/dockercourse:ch11_service_hw_start_failure \
hello-world

# run no-healtcheck container
docker service update \
 --update-failure-action rollback \
 --image omerurhan/dockercourse:ch11_service_hw_no_health \
 hello-world

# start healtcheck again
docker service update \
 --health-cmd /bin/httpping \
 --health-interval 10s \
 hello-world

# delete service
docker service rm hello-world
````

## Declarative service environments with Compose V3

### Collections of services with Compose V3
````
# run hello-world with stack
docker stack deploy -c hello-world.yml my-stack

docker stack ls
docker stack rm my-stack

# create multiple database app and display service discovery
docker stack deploy -c database.yml database
docker service ls
docker stack ps database

# update adminer to 3 replicas and update stack
docker stack deploy -c database.yml database

# display stack status
docker stack ps \
 --format '{{.Name}}\t{{.CurrentState}}' \
 database


# delete mariadb service with two method
docker stack deploy -c database.yml --prune database
docker stack deploy -c database.yml database && docker service rm mariadb

# delete stack
docker stack rm database
````
## Stateful services and preserving data

````
# deploy stack with volume
docker stack deploy \
-c database.yml \
--prune \
database

# list volumes
docker volume ls

# open adminer crete schema, table and delete service
docker service remove database_postgres
docker stack deploy \
-c database-with-volume.yml \
--prune \
database

docker stack rm database
docker volume ls
````

## Load balancing, service discovery, and networks with Compose

````
# inspect ingress network
docker network inspect ingress
docker network inspect database

docker stack deploy \
-c database-with-network.yml \
--prune \
database

# to use current network other stack use foo with  external: true property
````


