# Image Pipelines

## Patterns for building images

### All-in One Image

````
# build, run, inspect all-in one image 
docker build -t omerurhan/dockercourse:all-in-one -f all-in-one.df demo
docker run -d -p 8080:8080 omerurhan/dockercourse:all-in-one

docker image ls 
````
### Separate build and runtime 

````
cd demo

docker container run -it --rm \
-v "$(pwd)":/project/ \
-w /project/ \
maven:3.9.9-eclipse-temurin-21 \
mvn clean install

ls -la target/
cd ../
docker image build -t omerurhan/dockercourse:simple-runtime \
--file simple-runtime.df demo
docker container run -d -p 8080:8080 omerurhan/dockercourse:simple-runtime

docker image ls
````
## Record metadata at image build time

### Orchestrating the build with make
````
# make practices
make clean
make metadata
make app-artifacts
make lint-dockerfile
make app-image-simple-runtime
make app-image-all-in-one
make image-tests-simple-runtime
make inspect-image-labels-simple-runtime
````

## Testing images in a build pipeline

````
# run command and fileexisting test for simple-runtime container
make image-tests-simple-runtime BUILD_ID=$BUILD_ID
````

