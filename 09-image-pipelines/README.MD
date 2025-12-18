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
maven:3.3-jdk-8 \
mvn clean install

ls -la target/
docker image build -t omerurhan/dockercourse:simple-runtime \
--file simple-runtime.df .
docker container run -d -p 8080:8080 omerurhan/dockercourse:simple-runtime

docker image ls
````

