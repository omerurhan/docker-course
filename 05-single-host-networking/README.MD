# Single-Host Networking

This document explains common Docker networking tasks on a single host: creating networks, attaching containers, publishing ports, and basic network customizations.

Prerequisites
- Docker Engine running on your machine
- Basic familiarity with running containers and the command line

## Inspect networks

List available Docker networks:

```bash
docker network ls
```

## Create and use bridge networks

Bridge networks connect containers on the same Docker host. Use them when containers should see each other using internal IPs.

```bash
# Create a user bridge network with a custom subnet
docker network create \
	--driver bridge \
	--label project=dockercourse \
	--label chapter=5 \
	--attachable \
	--scope local \
	--subnet 10.0.42.0/24 \
	--ip-range 10.0.42.128/25 \
	user-network

# Run a temporary container attached to that network
docker run -it --rm --network user-network --name network-explorer alpine:3.8 sh

# Inside the container: show IPv4 addresses
ip -f inet -4 -o addr
```

To connect the same container to a second network from the host:

```bash
# Create a second bridge network
docker network create \
	--driver bridge \
	--label project=dockercourse \
	--label chapter=5 \
	--attachable \
	--scope local \
	--subnet 10.0.43.0/24 \
	--ip-range 10.0.43.128/25 \
	user-network2

# Attach the running container to the second network
docker network connect user-network2 network-explorer

# Back inside the container: check interfaces
ip addr

# (Optional) install nmap and probe both subnets
apk update && apk add --no-cache nmap
nmap -sn 10.0.42.0/24 -oG /dev/stdout | grep Status || true
nmap -sn 10.0.43.0/24 -oG /dev/stdout | grep Status || true
```

## Special network modes: host and none

- `host`: container shares the host network stack (Linux only).
- `none`: container has no external network interfaces.

```bash
# Host mode (Linux only): container uses host interfaces
docker run --rm --network host alpine:3.8 ip -o addr

# None: no network interfaces other than loopback
docker run --rm --network none alpine:3.8 ip -o addr

# Connectivity test in none mode (expected to fail)
docker run --rm --network none alpine:3.8 ping -w 2 1.1.1.1 || true
```

Note: `--network host` behaves differently on macOS/Windows because Docker runs inside a VM.

## Publishing ports (host ↔ container)

Map container ports to host ports with `-p`.

```bash
# UDP mapping: host UDP 8088 -> container UDP 8080
docker run --rm -p 8088:8080/udp alpine:3.8 echo "host UDP 8088 -> container UDP 8080"

# Bind ports to localhost only
docker run --rm -p 127.0.0.1:8080:8080/tcp -p 127.0.0.1:3000:3000/tcp alpine:3.8 echo "ports bound to localhost"

# Start a service and query its published ports
docker run -d -p 8080:80 --name listener nginx:latest
docker port listener
```

## DNS, hosts and other network customizations

You can override DNS servers, search domains and static host entries per container.

```bash
# Use a specific DNS server
docker run --rm --dns 8.8.8.8 alpine:3.8 nslookup docker.com

# DNS search and multiple DNS servers
docker run --rm --dns-search docker.com --dns 1.1.1.1 alpine:3.8 cat /etc/resolv.conf

# Add a static /etc/hosts entry
docker run --rm --add-host test:10.10.10.255 alpine:3.8 nslookup test

# Set container hostname and additional hosts
docker run --rm --hostname mycontainer \
	--add-host docker.com:127.0.0.1 --add-host test:10.10.10.2 alpine:3.8 cat /etc/hosts
```

Troubleshooting tips
- On macOS/Windows be aware of Docker Desktop's VM layer and how it affects networking.
- If a host port is already in use, Docker will fail to bind that port.
- Use `docker inspect`, `docker logs`, and `docker network inspect` to debug networking issues.

Run the commands interactively to learn how Docker networking behaves in your environment and adapt subnets and ports as needed.
