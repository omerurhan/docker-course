# Public and Private Software Distrubition

## Publishing with hosted registries

````
# build and push
docker image build \
 -t omerurhan/dockercourse:hello-dockerfile \
 -f HelloWorld.df \

docker login -u <username>
docker image push omerurhan/dockercourse:hello-dockerfile

# search
docker search omerurhan

````
## Introducing private registries

````
# run self-hosted docker registry
docker run -d -p 5000:5000 \
 -v "$(pwd)"/data:/tmp/registry-dev \
 --restart=always --name local-registry registry:2

# Using the registry image
docker image pull nginx

docker image tag nginx \
 localhost:5000/dockercourse/nginx

docker image push localhost:5000/dockercourse/nginx

# Consuming images from your registry
docker image rm \
 nginx \
 localhost:5000/dockercourse/nginx 

docker image ls
docker image pull localhost:5000/dockercourse/nginx 
docker image ls
````

## Manual image publishing and distribution


### A sample distribution infrastructure using FTP

````
docker image pull nginx:latest

docker network create ch9_ftp

docker container run -d --name ftp-server --network=ch9_ftp -p 21:21 \
omerurhan/dockercourse:ch9_ftpd

# upload image tar via ftp client
docker image save -o ./nginx.tar nginx:latest
docker container run --rm -it --network ch9_ftp \
 -v "$(pwd)":/data \
 omerurhan/dockercourse:ch9_ftp_client \
 -e 'cd pub/incoming; put nginx.tar; exit' ftp-server

# verify uploaded image
docker run --rm -it --network ch9_ftp \
 -v "$(pwd)":/data \
 omerurhan/dockercourse:ch9_ftp_client \
 -e "cd pub/incoming; ls; exit" ftp-server

#  First, eliminate the registry image from your local image cache and the file from your local directory:
rm nginx.tar
docker image rm nginx:latest
docker image ls

# Then download the image file from your FTP server by using the FTP client:
docker container run --rm -it --network ch9_ftp \
-v "$(pwd)":/data \
omerurhan/dockercourse:ch9_ftp_client \
-e 'cd pub/incoming; get nginx.tar; exit' ftp-server

# You can reload that image into your local cache with the docker load command:
docker image load -i nginx.tar
````


