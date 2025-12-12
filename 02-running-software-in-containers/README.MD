# Running Software in Containers — Practical Exercises

This document contains practical exercises for learning how to run and control software inside Docker containers. Each section includes short explanations and runnable commands you can use to experiment locally or in an online sandbox such as Play with Docker.

## Controlling containers: Building a website monitor

In this example you use three containers:

- `web`: an NGINX web server (runs detached/background)
- `mailer`: a small mailer service (runs detached/background)
- `agent`: an interactive watcher/monitor that inspects the other two

The goals of this exercise:

- Create detached and interactive containers
- List and inspect containers
- View and follow container logs
- Stop, restart, and remove containers
- Attach and detach terminals from containers

Note: `--link` is used here for simplicity in a workshop environment; it is a legacy feature and not recommended for production networks. Use user-defined networks for real projects.

### Start the services

```bash
# Start the web server (NGINX) in detached mode
docker run --detach \
  --name web \
  nginx:latest

# Start the mailer service in detached mode (image provided for this course)
docker run -d \
  --name mailer \
  omerurhan/dockercourse:ch2_mailer
```

### Run interactive tests and the agent

```bash
# Run a small BusyBox container interactively and link to the web container
docker run -it --link web:web --name web_test busybox:1.29 /bin/sh

# Inside the BusyBox shell you can test the web server:
wget -O - http://web:80/

# To detach from an attached container without stopping it press:
# Ctrl-p Ctrl-q

# Start the agent (interactive watcher) and link it to web and mailer
docker run -it \
  --name agent \
  --link web:insideweb \
  --link mailer:insidemailer \
  omerurhan/dockercourse:ch2_agent

# Detach from the agent container with: Ctrl-p Ctrl-q
```

### Inspect and control containers

```bash
# List running containers
docker ps

# Restart containers if needed
docker restart web
docker restart mailer
docker restart agent

# View logs (use -f to follow)
docker logs web
docker logs mailer
docker logs -f agent

# Stop and then inspect logs
docker stop web
docker logs mailer

# Remove containers when finished
docker rm web
docker rm agent
docker rm -f mailer
```

## PID namespaces: practical examples and troubleshooting

This section demonstrates PID namespaces and how processes are visible inside containers.

```bash
# Start two simple containers for namespace comparison
docker run -d --name namespaceA busybox:1.29 /bin/sh -c "sleep 30000"
docker run -d --name namespaceB busybox:1.29 /bin/sh -c "nc -l 0.0.0.0 -p 80"

# Inspect running processes inside each container
docker exec namespaceA ps
docker exec namespaceB ps

# To see host processes from a container (use with care):
docker run --pid host busybox:1.29 ps
```

## Container states: create, start, and inspect

You can create containers without starting them, inspect their state, then start only those you need.

```bash
# Create and run containers, capturing their IDs
MAILER_CID=$(docker run -d omerurhan/dockercourse:ch2_mailer)
WEB_CID=$(docker create nginx)
AGENT_CID=$(docker create --link $WEB_CID:insideweb \
  --link $MAILER_CID:insidemailer \
  omerurhan/dockercourse:ch2_agent)

# Show running and all containers
docker ps
docker ps -a

# Start containers by ID
docker start $AGENT_CID
docker start $WEB_CID
```

## Building environment-agnostic containers

Design containers so they behave the same regardless of the host. Examples below show using a read-only filesystem, writable temporary filesystems, and inspecting filesystem changes.

```bash
# Run WordPress with a read-only container filesystem
docker run -d --name wp --read-only wordpress:5.0.0-php7.2-apache

# Check whether the container is running and view logs
docker inspect --format "{{.State.Running}}" wp
docker logs wp

# Run a normal (writable) WordPress container for comparison
docker run -d --name wp_writable wordpress:5.0.0-php7.2-apache

# Show filesystem changes created by the writable container
docker diff wp_writable

# Run WordPress read-only but provide writable runtime areas (tmpfs)
docker run -d --name wp2 \
  --read-only \
  --tmpfs /tmp \
  wordpress:5.0.0-php7.2-apache

docker inspect --format "{{.State.Running}}" wp2
docker logs wp2
```

### Environment variable injection

Pass environment variables into containers at runtime using `-e` / `--env`.

```bash
docker run --env MY_ENVIRONMENT_VAR="this is a test" \
  busybox:1.29 env

# Example: start a MySQL container with a root password and the mailer
export DB_CID=$(docker run -d -e MYSQL_ROOT_PASSWORD=ch2demo mysql:5.7)
export MAILER_CID=$(docker run -d omerurhan/dockercourse:ch2_mailer)

# Run the multi-client starter script (project-specific)
CLIENT_ID=firstclient ./start-wp-multiple-clients.sh
```

## Building durable containers

Durable containers use restart policies or supervisors to remain available when processes fail.

```bash
# Run a simple container that always restarts on failure or host reboot
docker run -d --name backoff-detector --restart always busybox:1.29 date

# Follow logs and execute simple commands inside it
docker logs -f backoff-detector
docker exec backoff-detector echo "Just a test"

# Example of a container image that uses supervisord to keep multiple processes running
docker run -d -p 80:80 --name lamp-test omerurhan/dockercourse:lamp

docker top lamp-test
docker exec lamp-test ps

# If you intentionally kill a process inside the container, check logs and process table
docker exec lamp-test kill <PID>  # replace <PID> with the target process id
docker logs lamp-test

# Inspect the Docker entrypoint script shipped with the WordPress image
docker run --entrypoint="cat" wordpress:5.0.0-php7.2-apache /usr/local/bin/docker-entrypoint.sh
```

---

If you'd like, I can also:

- Convert these examples to use a user-defined Docker network instead of `--link`.
- Add step-by-step exercises with expected outputs and troubleshooting hints.
