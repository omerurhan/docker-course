
# Working with Storage and Volumes

This document contains practical examples for working with Docker bind mounts, in-memory mounts (tmpfs), and named volumes.

Note: On macOS the Docker engine runs inside a virtual machine; bind mount behavior can differ from Linux. Ensure host paths exist and have appropriate permissions before mounting.

## Bind mounts

Bind mounts map a directory or file from the host into a container. They are useful for sharing configuration files, logs, or other host resources with a container.

Example: run Nginx with a host configuration file and a host log file.

```bash
# Create a sample log file and Nginx config on the host
touch ~/example.log
cat > ~/example.conf <<'EOF'
server {
	listen 80;
	server_name localhost;
	access_log /var/log/nginx/custom.host.access.log main;
	location / {
		root /usr/share/nginx/html;
		index index.html index.htm;
	}
}
EOF

# Run Nginx using bind mounts for the config and the log file
CONF_SRC=~/example.conf
CONF_DST=/etc/nginx/conf.d/default.conf
LOG_SRC=~/example.log
LOG_DST=/var/log/nginx/custom.host.access.log

docker run -d --name diaweb \
	--mount type=bind,source=${CONF_SRC},target=${CONF_DST} \
	--mount type=bind,source=${LOG_SRC},target=${LOG_DST} \
	-p 80:80 \
	nginx:latest

# View container logs (note: Nginx writes to the mounted host file)
docker logs diaweb
cat ~/example.log

# Stop and remove the container
docker rm -f diaweb

# Run again with the config mount set to read-only on the host side
docker run -d --name diaweb \
	--mount type=bind,source=${CONF_SRC},target=${CONF_DST},readonly \
	--mount type=bind,source=${LOG_SRC},target=${LOG_DST} \
	-p 80:80 \
	nginx:latest

# Demonstrate editing inside the container (affects only container filesystem when config is read-only)
docker exec diaweb sed -i "s/listen 80/listen 8080/" /etc/nginx/conf.d/default.conf || true

# Cleanup
docker rm -f diaweb || true
```

Notes:
- If you mount a host file into a container, the host file must already exist.
- On macOS and Windows, file permission and ownership behavior may differ from Linux.

## In-memory storage (tmpfs)

Use tmpfs mounts for ephemeral, memory-backed storage. Data stored on tmpfs is not persisted to disk and is lost when the container stops.

```bash
docker run --rm \
	--mount type=tmpfs,destination=/tmp \
	--entrypoint mount \
	alpine:latest -v

# You should see a mount line for tmpfs (e.g. "tmpfs on /tmp type tmpfs")
```

## Docker volumes

Named volumes are managed by Docker and are the recommended way to persist data used by containers.

```bash
# Create a named volume
docker volume create --driver local --label example=location location-example

# List and inspect volumes
docker volume ls
docker volume inspect --format '{{json .Mountpoint}}' location-example
```

Example: using a named volume with Cassandra

```bash
# Create a volume for Cassandra data
docker volume create --label example=cassandra cass-shared

# Start a Cassandra container using the named volume
docker run -d --name cass2 \
	--mount type=volume,source=cass-shared,target=/var/lib/cassandra/data \
	cassandra:2.2

# Connect to Cassandra from a temporary client container
docker run -it --rm --link cass2:cass cassandra:2.2 cqlsh cass

# In cqlsh you can create a keyspace and query schema_keyspaces as shown in the exercises.

# When finished, remove containers and the volume
docker rm -vf cass2 || true
docker volume rm cass-shared || true
```

## Shared mount points and sharing files

This section shows two patterns for sharing files between the host and containers: bind mounts (host directory) and named volumes (managed by Docker).

Bind mount example (sharing logs created by a writer container):

```bash
LOG_SRC=~/web-logs-example
mkdir -p "${LOG_SRC}"

# Start a writer container that writes logs to /data inside the container
docker run --name plath -d \
	--mount type=bind,source=${LOG_SRC},target=/data \
	omerurhan/dockercourse:ch4_writer_a

# Inspect the first lines of the shared log from a lightweight container
docker run --rm --mount type=bind,source=${LOG_SRC},target=/data alpine:latest head /data/logA || true

# Inspect the file on the host
cat ${LOG_SRC}/logA || true

docker rm -f plath || true
```

Named volume example (same writer, using a Docker volume):

```bash
# Create a named volume
docker volume create --driver local logging-example

# Start writer container using the named volume
docker run --name plath -d --mount type=volume,source=logging-example,target=/data omerurhan/dockercourse:ch4_writer_a

# Read the data from the volume using a temporary container
docker run --rm --mount type=volume,source=logging-example,target=/data alpine:latest head /data/logA || true

# Find where the volume is stored on the host (Linux Docker only)
docker volume inspect --format '{{json .Mountpoint}}' logging-example

# On Linux you can also inspect the underlying files directly (path printed above):
# sudo cat /var/lib/docker/volumes/logging-example/_data/logA

docker stop plath || true
docker rm -f plath || true
docker volume rm logging-example || true
```

## Cleanup and tips

- Remove unused volumes: `docker volume prune`
- Remove stopped containers: `docker container prune`
- If you see permission issues when using bind mounts, check host file ownership and Docker VM settings (macOS/Windows)

This file provides focused examples — run the commands interactively to observe behavior and experiment with mount options and permissions.
```
