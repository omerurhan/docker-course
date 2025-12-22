# Abstraction of Configuration

## Configuration distribution and management

````
# deploy application for dev env with dev config.
DEPLOY_ENV=dev docker stack deploy \
 --compose-file docker-compose.yml greetings_dev

# Managing config resources directly 
docker config inspect greetings_dev_env_specific_config 

docker service inspect \
--format '{{ json .Spec.TaskTemplate.ContainerSpec.Configs }}' \
greetings_dev_api

# remove stack
docker stack rm greetings_dev

# deploy stage
DEPLOY_ENV=stage docker stack deploy \
 --compose-file docker-compose.yml greetings_dev

# remove stack
docker stack rm greetings_stage
````
## Secrets - A special kind of configuration

````
# deploy secure greetings app
DEPLOY_ENV=prod docker stack deploy --compose-file docker-compose.yml \
--compose-file docker-compose.prod.yml \
greetings_prod

# define secret 
cat api/config/insecure.key | \
 docker secret create ch12_greetings-svc-prod-TLS_PRIVATE_KEY_V1 -

# inspect secret
docker secret inspect ch12_greetings-svc-prod-TLS_PRIVATE_KEY_V1

# create certificate config
docker config create \
 ch12_greetings_svc-prod-TLS_CERT_V1 api/config/insecure.crt

# deploy secure greetings app again
DEPLOY_ENV=prod docker stack deploy --compose-file docker-compose.yml \
--compose-file docker-compose.prod.yml \
greetings_prod

# verify
docker service ps greetings_prod_api

# display log
docker service logs --since 2m greetings_prod_api
````
