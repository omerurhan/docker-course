# Docker Training Course

Welcome to the Docker Training Course repository — a compact, hands-on set of lessons and examples that introduce container fundamentals and demonstrate how to run real software inside containers.

**Course Overview**

This course is organized as a sequence of focused lessons (01–12). Each lesson contains short explanations and runnable examples so you can practice concepts immediately.

**Project Structure and Lesson Descriptions**

- `01-welcome-to-docker/`: Introductory material explaining container fundamentals, images, containers, and basic Docker commands.
- `02-running-software-in-containers/`: Examples and exercises for running real applications inside containers, including a sample script to start a WordPress multi-container setup.
- `03-software-installation-simplified/`: Dockerfile examples and helper scripts demonstrating how container images can simplify software installation and packaging.
- `04-working-with-storage-and-volumes/`: Lessons and examples focused on persistent storage, volumes, and data management for containers.
- `05-single-host-networking/`: Networking basics on a single Docker host, including how containers communicate and how to expose services.
- `06-limiting-risk-with-resource-control/`: Examples showing resource constraints (CPU, memory) and how to limit container resource usage.
- `07-packaging-software-in-images/`: Multiple Dockerfile examples, multistage builds, and utilities for creating efficient, production-ready images.
- `08-public-and-private-software-distrubition/`: Demonstrates workflows for sharing images publicly and privately, plus best practices for distribution.
- `09-image-pipelines/`: CI/CD and image pipeline examples, including multi-stage build patterns and structure tests for image validation.
- `10-services-with-docker-and-compose/`: `docker-compose` examples that define multi-container services, databases, and simple application stacks.
- `11-abstraction-of-configuration/`: Configuration management patterns with environment-specific Compose files and examples that show runtime configuration handling.
- `12-cluster-of-docker-hosts-with-docker-swarm/`: Intro to Docker Swarm for managing clusters, service deployment, and simple orchestration examples.

**Course Goals**

By working through this repository you will learn how to:

- Build and run containers from Dockerfiles
- Work with container images and registries
- Use Docker Compose for multi-container applications
- Manage storage and networking for containers
- Apply resource constraints and secure runtime configurations
- Package, test, and distribute images for different environments

**Prerequisites**

- Docker Desktop (macOS) or Docker Engine and Docker Compose installed
- Basic command-line familiarity (zsh / bash)

**Quick Start**

Open a terminal and run these commands to explore the repository and try an example:

```bash
# List repository contents
ls -la

# Read the introductory lesson
less 01-welcome-to-docker/README.MD

# Example: run the WordPress multi-client script (make executable if needed)
cd 02-running-software-in-containers
chmod +x start-wp-multiple-client.sh
./start-wp-multiple-client.sh
```

If a script requires `docker-compose` or elevated privileges, run it with an account that has Docker access or prefix commands with `sudo` as needed.

**How to Use This Course**

- Start with lesson `01-welcome-to-docker` to learn core concepts.
- Follow the lessons in order or jump to topics of interest (networking, storage, packaging).
- Inspect Dockerfiles and scripts, then run the provided examples to see practical results.

**Contributing & Notes**

- Examples are intentionally small and focused for learning and experimentation.
- If you improve examples or add clarifications, update the relevant lesson README so others benefit.

**Getting Help**

If you get stuck, open an issue in your project's issue tracker or consult Docker's documentation: https://docs.docker.com/

---

Happy learning — containers make deployments reproducible and development environments portable. Start with the introductory lesson and then run the examples to gain hands-on experience.