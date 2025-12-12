# Docker Training Course

Welcome to the Docker training repository — a compact, hands-on course that introduces container fundamentals and shows how to run real software inside containers.

**Project Structure**

- `01-welcome-to-docker/`
  - `01-welcome-to-docker/README.MD` — Beginner-friendly introduction to Docker concepts, images, containers, and basic commands.

- `02-running-software-in-containers/`
  - `02-running-software-in-containers/README.MD` — Hands-on exercises and examples for running real applications in containers.
  - `02-running-software-in-containers/start-wp-multiple-client.sh` — Example script to start a WordPress environment (multi-client/sample setup).

**Course Goal**

This repository is designed to help learners move from zero to practical Docker skills by combining short explanations with runnable examples. You will learn how to:

- Build and run containers
- Use Docker images and registries
- Compose and orchestrate simple multi-container apps
- Debug and inspect running containers

**Prerequisites**

- Docker Desktop (macOS) or Docker Engine and Docker Compose installed
- Basic command-line familiarity (zsh / bash)

**Quick Start**

Open a terminal and try the following commands to explore and run the provided example:

```bash
# List repository contents
ls -la

# Read the introductory material
less 01-welcome-to-docker/README.MD

# Run the example WordPress script (example assumes executable permission)
cd 02-running-software-in-containers
chmod +x start-wp-multiple-client.sh
./start-wp-multiple-client.sh
```

If the script requires `docker-compose` or `docker` access, you may need to run it as a user with Docker permissions or prefix with `sudo` depending on your local setup.

**How to Use This Course**

- Start with `01-welcome-to-docker/README.MD` to learn the core concepts.
- Move to `02-running-software-in-containers/README.MD` for practical, runnable examples.
- Inspect and run the `start-wp-multiple-client.sh` script to see a small multi-container application in action.

**Further notes**

- Files and scripts are intentionally small and focused for learning and experimentation.
- If you make changes or improvements, consider adding notes to the subfolder README files so other learners can follow your work.

**Getting Help**

If you get stuck, open an issue in your project tracker or consult Docker's official docs: https://docs.docker.com/

---

Happy learning — containers make deployments reproducible and development environments portable. Start with the welcome lesson and then run the example to see Docker in action.