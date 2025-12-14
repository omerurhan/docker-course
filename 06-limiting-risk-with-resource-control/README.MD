# Limiting Risk with Resource Control

This chapter shows how to limit a container's resource usage, isolate shared memory, manage user identity inside images, and restrict OS features using Linux capabilities and privilege flags.

## Resource limits (memory and CPU)

Use Docker's run flags to limit memory and CPU usage for containers. These controls help protect the host and other containers from resource exhaustion.

Examples:

```bash
# Start MariaDB with a memory limit and relative CPU weight
docker run -d --name ch6_mariadb \
	--memory 256m \
	--cpu-shares 1024 \
	--cap-drop net_raw \
	-e MYSQL_ROOT_PASSWORD=test \
	mariadb:5.5

# Monitor resource usage
docker stats ch6_mariadb

# Start WordPress with a memory limit and CPU shares
docker run -d -P --name ch6_wordpress \
	--memory 512m \
	--cpu-shares 512 \
	--cap-drop net_raw \
	--link ch6_mariadb:mysql \
	-e WORDPRESS_DB_PASSWORD=test \
	wordpress:5.0.0-php7.2-apache

# Limit CPU quota via CFS (Completely Fair Scheduler)
docker run -d -P --name ch6_wordpress_quota \
	--memory 512m \
	--cpus 0.75 \
	--cap-drop net_raw \
	--link ch6_mariadb:mysql \
	-e WORDPRESS_DB_PASSWORD=test \
	wordpress:5.0.0-php7.2-apache
```

Notes:
- `--memory` enforces an upper bound on RAM usage.
- `--cpu-shares` sets a relative weight for CPU scheduling when CPU is contended.
- `--cpus` (or `--cpu-quota`/`--cpu-period`) sets a hard CPU quota.

## Sharing memory (IPC)

Containers can share inter-process communication namespaces to exchange data via shared memory (e.g. /dev/shm). Use the `--ipc` flag carefully — sharing IPC increases coupling between containers.

```bash
# Start a producer container with a shareable IPC namespace
docker run -d -u nobody --name ch6_ipc_producer --ipc shareable omerurhan/dockercourse:ch6_ipc -producer

# Run a consumer that joins the producer's IPC namespace
docker run -d --name ch6_ipc_consumer --ipc container:ch6_ipc_producer omerurhan/dockercourse:ch6_ipc -consumer

# View logs for both containers
docker logs ch6_ipc_producer
docker logs ch6_ipc_consumer

# Cleanup
docker rm -vf ch6_ipc_producer ch6_ipc_consumer
```

If a consumer is started without joining the producer IPC, it will not see the shared memory objects.

## User identity inside images

Containers run with a user configured in the image by default (often `root`). You can inspect image configuration and run containers as a specific user.

```bash
# Inspect the image's default user
docker image pull busybox:1.29
docker image inspect busybox:1.29 --format '{{.Config.User}}'

# Run commands as the image's default user
docker run --rm --entrypoint "" busybox:1.29 whoami
docker run --rm --entrypoint "" busybox:1.29 id

# List users defined in the image
docker run --rm busybox:1.29 awk -F: '{print $1}' /etc/passwd

# Run a container as an unprivileged user
docker run --rm --user nobody busybox:1.29 id
docker run --rm -u nobody:nogroup busybox:1.29 id
docker run --rm -u 10000:20000 busybox:1.29 id
```

Best practice: run services as non-root users inside the container when possible.

## Linux capabilities (adjusting OS feature access)

Linux capabilities let you grant or drop fine-grained kernel privileges without running a container as full `--privileged`. Use `--cap-drop` to remove, and `--cap-add` to add capabilities.

```bash
# Show capabilities available inside the container (requires libcap utilities inside image)
docker run --rm ubuntu:16.04 bash -c "apt-get update >/dev/null && apt-get install -y libcap2-bin >/dev/null && capsh --print | grep net_raw || true"

# Run without the net_raw capability
docker run --rm --cap-drop net_raw ubuntu:16.04 bash -c "apt-get update >/dev/null && apt-get install -y libcap2-bin >/dev/null && capsh --print | grep net_raw || true"

# Add CAP_SYS_ADMIN (powerful — avoid unless necessary)
docker run --rm --cap-add SYS_ADMIN ubuntu:16.04 bash -c "apt-get update >/dev/null && apt-get install -y libcap2-bin >/dev/null && capsh --print | grep sys_admin || true"
```

Notes:
- Dropping capabilities reduces an attack surface while keeping the container usable.
- `CAP_SYS_ADMIN` is effectively very powerful — avoid granting it unless required.

## Privileged containers

Running a container with `--privileged` gives it almost all host capabilities and device access. This is useful for debugging or running system-level tools, but it bypasses many security boundaries and should be avoided in production.

```bash
# Start a one-off privileged container for inspection
docker run --rm --privileged ubuntu:16.04 id

# Inspect kernel capabilities and devices (for debugging only)
docker run --rm --privileged ubuntu:16.04 capsh --print
docker run --rm --privileged ubuntu:16.04 ls /dev | head
docker run --rm --privileged ubuntu:16.04 networkctl || true
```

Security tips
- Prefer dropping capabilities (`--cap-drop`) and running as non-root over `--privileged`.
- Use resource limits to prevent noisy neighbors and DoS from containers.
- Test changes in a safe environment before applying them to production.

This file contains runnable examples — adjust image names, users, and limits to match your environment before executing.

