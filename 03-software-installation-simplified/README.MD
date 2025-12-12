# Software Installation Simplified

This section provides practical examples demonstrating how software is
defined and installed within Docker environments.

------------------------------------------------------------------------

## 1. Finding and Installing Software

### **Using Alternative Registries**

You can pull images from registries other than Docker Hub, such as Quay:

``` bash
docker pull quay.io/nginx/nginx-unprivileged
```

General registry reference format:

    [REGISTRYHOST:PORT/][USERNAME/]NAME[:TAG]

Remove the image:

``` bash
docker rmi quay.io/nginx/nginx-unprivileged
docker images
```

------------------------------------------------------------------------

### **Working with Images as Files**

Pull an image:

``` bash
docker pull busybox:latest
```

Save the image to a tar archive:

``` bash
docker save -o myfile.tar busybox:latest
```

Remove the image:

``` bash
docker rmi busybox
```

Load the image back from the tar file:

``` bash
docker load -i myfile.tar
```

------------------------------------------------------------------------

### **Installing from a Dockerfile**

Navigate to the directory containing the Dockerfile:

``` bash
cd 03-software-installation-simplified/dockerfile
```

Build the image:

``` bash
docker build -t omerurhan/dockercourse:dockerfile .
```

------------------------------------------------------------------------

## 2. Installation Files and Isolation

### **Image Layers in Action**

Pull the images (this only downloads them; it does not start
containers):

``` bash
docker pull dockerinaction/ch3_myapp
docker pull dockerinaction/ch3_myotherapp
```

------------------------------------------------------------------------

### **Understanding Layer Relationships**

Run a base container:

``` bash
docker run --name=base-container -ti ubuntu
```

Inside the container, install Node.js:

``` bash
apt update && apt install -y nodejs
node -e 'console.log("Hello world!")'
```

Commit the modified container into a new image:

``` bash
docker container commit -m "Add node" base-container node-base
```

Inspect the image layers:

``` bash
docker image history node-base
```

Run a container from the new image to confirm Node.js is installed:

``` bash
docker run node-base node -e "console.log('Hello again')"
```
