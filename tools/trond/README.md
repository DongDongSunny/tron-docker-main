# TROND

This document provides guidance on using the `trond` command-line tool, tested on Linux and macOS.

## Prerequisites

- Shell (for running the sh scripts)

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/tronprotocol/tron-docker.git
    cd tron-docker/
    ```

2. Build the `trond` binary:

    ```sh
    ./build_trond.sh
    ```

## Usage

The [trond](tools/trond/docs/trond.md) command-line tool offers various commands for managing TRON nodes and snapshots.

### Examples

- Download the latest mainnet lite fullnode snapshot from the default source (`34.143.247.77`) to the current directory:

    ```sh
    ./trond snapshot download default-main
    ```

- Run a single java-tron fullnode for the main network:

    ```sh
    ./trond node run-single -t full-main
    ```

For more information, refer to the [trond documentation](tools/trond/docs/trond.md).

## Note

This tool is in an early phase and has only been tested on macOS and Linux. If you encounter any issues, raise them here [GitHub](https://github.com/tronprotocol/tron-docker/issues).