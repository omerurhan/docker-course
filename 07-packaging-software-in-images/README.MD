# Packaging Software in Images


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

### Metadata instructions

````
# please atteion for .dockerignore file
docker image build -t omerurhan/dockercouse:0.6 -f Dockerfile.metadata .
docker inspect omerurhan/dockercouse:0.6
````

