# Cluster of Docker Hosts with Docker Swarm


## Deploying Application to Swarm Cluster

````
# list nodes on worker and manager
docker node ls

# deploy app on manager node
docker stack deploy --compose-file docker-compose.yml multi-tier-app

# create secret
echo 'mydbpass72' | docker secret create \
 ch13_multi_tier_app-POSTGRES_PASSWORD -

# display secret infos 
docker secret ls --format "table {{.ID}} {{.Name}} {{.CreatedAt}}" 

# deploy again
docker stack deploy \
 --compose-file docker-compose.yml multi-tier-app

# display services
docker service ls \
 --format "table {{.Name}} {{.Mode}} {{.Replicas}}" 

# display logs
docker service logs --follow multi-tier-app_api

# send request to app
curl http://localhost:8080

# display responses
for i in `seq 1 4`; do curl http://localhost:8080; sleep 1; done; 
````

## Communicating with services running on a Swarm cluster

````
# inspect services
docker service inspect --format="{{json .Endpoint.Spec.Ports}}" \
 multi-tier-app_api

# display virtual ips
docker service inspect --format '{{ json .Endpoint.VirtualIPs }}' \
 multi-tier-app_api
````

### Isolating service-to-service communication with overlay networks

````
# create new service
docker service create --name echo --publish '8000:8' busybox:1.29 \
 nc -v -lk -p 8 -e /bin/cat

# check from host
echo "Hello netcat my old friend, I've come to test connections again." \
 | nc -v -w 3 <your-host-ip> 8000

# check from another network
docker container run --rm -it --network multi-tier-app_private \
 alpine:3.8 sh

ping -c 1 echo
nslookup echo
````

## Placing Service Task on Cluster

````
# drain manager nodes
 docker node update --availability drain manager1
 docker node update --availability drain manager2
 docker node update --availability drain manager3

# display availability
docker node ls --format 'table {{ .ID }} {{ .Hostname }} {{ .Availability }}'

# verify migrated apps
docker service ps multi-tier-app_api multi-tier-app_postgres \
--filter 'desired-state=running' \
--format 'table {{ .Name }} {{ .Node }}'

# placement constraints with label
docker node update --label-add zone=private worker1
docker node update --label-add zone=public worker2

# relocate app with using constraints
docker service update \
--constraint-add 'node.labels.zone == public' \
multi-tier-app_api

# display multi-tier app
docker service ps multi-tier-app_api \
--filter 'desired-state=running' \
--format 'table {{ .Name }} {{ .Node }}'

# display node hostname, role, labels
for node_id in $(docker node ls -q | head); do
docker node inspect \
--format '{{.Description.Hostname}} {{.Spec.Role}} {{.Spec.Labels}}' \
 "${node_id}";
done;

# add record to postgre
curl http://127.0.0.1:8080/counter

# relocate postgres to private zone
docker service update --constraint-add 'node.labels.zone == private' \
 multi-tier-app_postgres

# verify relocation 
docker service ps multi-tier-app_postgres \
 --filter 'desired-state=running' \
 --format 'table {{ .Name }} {{ .Node }}'

# add new record to show data not migrated between nodes
curl http://127.0.0.1:8080/counter

# run postgresql dedicated node
docker service update --constraint-add 'node.hostname == worker1' \
 multi-tier-app_postgres
````

### Using global services for one task per node

````
# deploy one container per available node (global)
docker service create --name echo-global \
 --mode global \
 --publish '9000:8' \
 busybox:1.29 nc -v -lk -p 8 -e /bin/cat

# display pods 
docker service ps echo-global \
 --filter 'desired-state=running' \
 --format 'table {{ .Name }} {{ .Node }}'
````



