# Welcome to Docker

This section covers beginner-level Docker applications and practical examples to help you get started with containerization. 

## Prerequisites

Before starting, ensure you have the following:

- **Docker Hub Account**: Sign up at [Docker Hub](https://hub.docker.com/) to access Docker images
- **Play with Docker Access**: Have internet access to [Play with Docker](https://labs.play-with-docker.com/) for hands-on practice


## Common Docker Commands

Here are some essential Docker commands you'll use frequently:

### Running Your First Container

```bash
# Run the hello-world image
docker run hello-world

# Run the same command again
docker run hello-world
```

### Getting Help

```bash
# View all available Docker commands
docker help

# Get help for a specific command (e.g., copy files)
docker help cp

# Explore additional Docker CLI commands
docker --help
```

### Useful Tips

- Each `docker run` command creates a new container
- Use `docker help <command>` to learn more about any command
- Visit the [Docker Documentation](https://docs.docker.com/) for detailed guides
