# TRON Docker

This repository provides tools and guidance to help the community quickly get started with the TRON network and development.

## Features

### 🚀 Quick start for single FullNode
Easily deploy a single TRON FullNode connected to **Mainnet** or **Nile testnet** using Docker. Follow the instructions to get your node up and running in minutes.

### 🔗 Private chain setup
Set up your own private TRON blockchain network for development and testing. The provided configurations simplify deployment and management, making it ideal for custom use cases.

### 📊 Node monitoring with Prometheus and Grafana
Monitor the health and performance of your TRON nodes with integrated **Prometheus** and **Grafana** services. Real-time metrics and visualizations are just a few steps away.

### 🛠️ Tools
We also provide tools to facilitate the CI and testing process:
- **Gradle Docker**: Automate the build and testing of the `java-tron` Docker image using Gradle.
- **DBFork**: Launch a private java-tron network based on the Mainnet database state, enabling shadow fork testing.

## Prerequisites
Please download and install the latest version of Docker from the official Docker website:
* Docker Installation for [Mac](https://docs.docker.com/docker-for-mac/install/)
* Docker Installation for [Windows](https://docs.docker.com/docker-for-windows/install/)
* Docker Installation for [Linux](https://docs.docker.com/desktop/setup/install/linux/)


## Quick start
To quickly start a fullnode that connect to the Mainnet, simply download and run [docker-compose-quick-start.yml](single_node/docker-compose-quick-start.yml):
```
docker-compose -f docker-compose-quick-start.yml up
```
Once the fullnode starts, it will begin to sync blocks with Mainnet from genesis block.

You could use `docker exec tron-node tail -f ./logs/tron.log` to check the logs. For more details please refer [single_node](single_node/README.md) part.

## Start all services
First clone the repository:

```sh
git clone https://github.com/tronprotocol/tron-docker.git
cd tron-docker
```

### Start easily by Trond script
We provide Trond script to help developer easily start features using one command, enabling the community to quickly get started with TRON network development and interaction.

```
./trond -h
Usage:
  trond [command]

Examples:
# Help information for java-tron docker image build and testing command
$ ./trond docker

# Help information for database snapshot download related command
$ ./trond snapshot

# Help information for Tron single FullNode deployment command
$ ./trond node

Available Commands:
  docker      Commands for operating java-tron docker image.
  help        Help about any command
  node        Commands for operating java-tron docker node.
  snapshot    Commands for getting java-tron node snapshots.
```
Now it support features:
- Deploy single FullNode for varies networks.
- java-tron docker image build and testing.
- Download database snapshot.

### Start the service individually
To start all available features, navigate to the appropriate directory and follow the instructions in the corresponding README to start the services:
- **TRON network deployment related:**
   - **Single FullNode**: Use the [single_node](./single_node) folder.
   - **Private TRON network**: Use the [private_net](./private_net) folder.
   - **Node monitoring**: Use the [metric_monitor](./metric_monitor) folder.

- **Tools**:
   - **Gradle Docker**: Automate Docker image builds and testing. Check the [gradle docker](./tools/docker/README.md) documentation.
   - **DBFork**: Perform shadow fork testing. Follow the [DBFork guidance](./tools/dbfork/README.md).

## Troubleshooting
If you encounter any difficulties, please refer to the [Issue Work Flow](https://tronprotocol.github.io/documentation-en/developers/issue-workflow/#issue-work-flow), then raise an issue on [GitHub](https://github.com/tronprotocol/tron-docker/issues). For general questions, please use [Discord](https://discord.gg/cGKSsRVCGm) or [Telegram](https://t.me/TronOfficialDevelopersGroupEn).

## Contributing

All contributions are welcome. Check [contribution](CONTRIBUTING.md) for more details.

## License

This repository is released under the [LGPLv3 license](https://github.com/tronprotocol/tron-docker/blob/main/LICENSE).
