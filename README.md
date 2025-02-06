# TRON Docker

This repository provides tools and guidance to help the community quickly get started with the TRON network and development.

---

## Features

### 🚀 Quick Start for Single FullNode
Easily deploy a single TRON FullNode connected to **Mainnet** or **Nile testnet** using Docker. Follow the instructions to get your node up and running in minutes.

### 🔗 Private Chain Setup
Set up your own private TRON blockchain network for development and testing. The provided configurations simplify deployment and management, making it ideal for custom use cases.

### 📊 Node Monitoring with Prometheus and Grafana
Monitor the health and performance of your TRON nodes with integrated **Prometheus** and **Grafana** services. Real-time metrics and visualizations are just a few steps away.

### 🛠️ Tools
We also provide tools to facilitate the CI and testing process:
- **Gradle Docker**: Automate the build and testing of the `java-tron` Docker image using Gradle.
- **DB Fork**: Launch a private java-tron network based on the Mainnet database state, enabling shadow fork testing.

---

## Getting Started

### Prerequisites
- Docker
- Docker Compose

### Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/tronprotocol/tron-docker.git
   cd tron-docker
   ```

2. **Start the services:**
   Navigate to the relevant directory and follow the instructions in the respective README to start the services.

   **For TRON Network Deployment Related:**
   - **Single FullNode**: Use the [single_node](./single_node) folder.
   - **Private TRON Network**: Use the [private_net](./private_net) folder.
   - **Node Monitoring**: Use the [metric_monitor](./metric_monitor) folder.

   **For Tools:**
   - **Gradle Docker**: Automate Docker image builds and testing. Check the [gradle docker](./tools/docker/README.md) documentation.
   - **DB Fork**: Perform shadow fork testing. Follow the [dbfork guidance](./tools/dbfork/README.md).

## Troubleshooting
If you encounter any difficulties, please refer to the [Issue Work Flow](https://tronprotocol.github.io/documentation-en/developers/issue-workflow/#issue-work-flow), then raise an issue on [GitHub](https://github.com/tronprotocol/tron-docker/issues). For general questions, please use [Discord](https://discord.gg/cGKSsRVCGm) or [Telegram](https://t.me/TronOfficialDevelopersGroupEn).

# Contributing

All contributions are welcome. Check [contribution](CONTRIBUTING.md) for more details.

# License

This repository is released under the [LGPLv3 license](https://github.com/tronprotocol/tron-docker/blob/main/LICENSE).
