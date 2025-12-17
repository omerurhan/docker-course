# Packaging Software in Images


## Packaging with Dockerfile
````
# build docker image from Dockerfile
docker image build --tag ubuntu-git:auto .
docker image ls

# run container
docker container run --rm ubuntu-git:auto

# run build again to observe cache, if you don't want use cache you use  --no-cache parameter. 
docker image build --tag ubuntu-git:auto .
docker image build --no-cache --tag ubuntu-git:auto .

# failed docker build
docker image build -f Dockerfile.failed --tag ubuntu-git:auto .
````

### Metadata instructions

````
# please attention for .dockerignore file 
docker image build -t omerurhan/dockercouse:0.6 -f metadata.df .
docker inspect omerurhan/dockercouse:0.6

````
### Filesystem instructions
````
#build mailer dockerfile to display filesystem instructions
docker image build -t omerurhan/dockercourse:mailer-logging -f mailer-logging.df .

## run, inspect and test container
docker run -d --name logging-mailer -p 33333:33333 omerurhan/dockercourse:mailer-logging
docker port logging-mailer

## send dummy mail with netcat and check logs
nc localhost 33333
write something...
CTRL +C

docker exec -it logging-mailer cat /var/log/mailer.log
docker exec -it logging-mailer ls -l /var/log/mailer.log
````
##  Creating maintainable Dockerfiles

````
# ARG usage
version=0.6; docker image build -t omerurhan/docker-course:${version} \
 -f metadata-with-arg.df \
 --build-arg VERSION=${version} .

version=0.7; docker image build -t omerurhan/docker-course:${version} \
 -f metadata-with-arg.df \
 --build-arg VERSION=${version} .

# multi-stage build 
docker build -t omerurhan/dockercourse:http-client -f http-client.df .

# display image size and run container
docker image ls
docker run omerurhan/dockercourse:http-client 
````

## Using Startup Scripts and Multiprocess Containers

````
# investigate example-startup.sh to understand how startup script working

# The purpose and use of health checks
docker image build -t omerurhan/dockercourse:healthcheck .
docker container run --name healthcheck_ex -d omerurhan/dockercourse:healthcheck

docker ps  // display healthcheck.
docker ps --format 'table {{.Names}}\t{{.Image}}\t{{.Status}}'

# override healtcheck if it exist
docker container run --name=healthcheck_ex -d \
--health-cmd='nc -vz -w 2 localhost 80 || exit 1' \
nginx:1.13-alpine

````
