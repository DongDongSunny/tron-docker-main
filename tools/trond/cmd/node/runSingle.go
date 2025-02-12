package node

import (
	"fmt"
	"log"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spf13/cobra"
	"github.com/tronprotocol/tron-docker/utils"
)

var runSingleCmd = &cobra.Command{
	Use:   "run-single",
	Short: "Run single java-tron node for different networks.",
	Long: heredoc.Doc(`
			You need to make sure the local environment is ready before running the node. Run "./trond node env" to check the environment before starting the node.

			The following files are required:

				- Database directory: ./output-directory
				- Configuration file(by default, these exist in the current repository directory)
					main network: ./conf/main_net_config.conf
					nile network: ./conf/nile_net_config.conf
					private network: ./conf/private_net_config_*.conf
				- Docker compose file(by default, these exist in the current repository directory)
					main network: ./single_node/docker-compose.fullnode.main.yml
					nile network: ./single_node/docker-compose.fullnode.nile.yml
					private network: ./single_node/docker-compose.witness.private.yml
				- Log directory: ./logs
		`),
	Example: heredoc.Doc(`
			# Run single java-tron fullnode for main network
			$ ./trond node run-single -t full-main

			# Run single java-tron fullnode for nile network
			$ ./trond node run-single -t full-nile

			# Run single java-tron witness node for private network
			$ ./trond node run-single -t witness-private
		`),
	Run: func(cmd *cobra.Command, args []string) {
		nType, _ := cmd.Flags().GetString("type")

		dockerComposeFile := ""
		switch nType {
		case "full-main":
			dockerComposeFile = "./single_node/docker-compose.fullnode.main.yml"
		case "full-nil":
			dockerComposeFile = "./single_node/docker-compose.fullnode.nile.yml"
		case "witness-private":
			dockerComposeFile = "./single_node/docker-compose.witness.private.yml"
		default:
			fmt.Println("Error: type not supported", nType)
		}
		if yes, isDir := utils.PathExists(dockerComposeFile); !yes || isDir {
			fmt.Println("Error: file not exists or not a file:", dockerComposeFile)
		}
		if err := utils.RunComposeServiceOnce(dockerComposeFile); err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Node started successfully.")
		fmt.Println("You can check the log file in ./logs directory. For example, run 'tail -f ./logs/tron-node.log' to check the log.")
		fmt.Println("You can also check the log in the docker container by running 'docker logs -f tron-node'.")
	},
}

var runSingleStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop single java-tron node for different networks.",
	Long: heredoc.Doc(`
			The following configuration files are required:

				- Configuration file(by default, these exist in the current repository directory)
					main network: ./conf/main_net_config.conf
					nile network: ./conf/nile_net_config.conf
					private network: ./conf/private_net_config_*.conf
				- Docker compose file(by default, these exist in the current repository directory)
					main network: ./single_node/docker-compose.fullnode.main.yml
					nile network: ./single_node/docker-compose.fullnode.nile.yml
					private network: ./single_node/docker-compose.witness.private.yml
		`),
	Example: heredoc.Doc(`
			# Run single java-tron fullnode for main network
			$ ./trond node run-single stop -t full-main

			# Run single java-tron fullnode for nile network
			$ ./trond node run-single stop -t full-nile

			# Run single java-tron witness node for private network
			$ ./trond node run-single stop -t witness-private
		`),
	Run: func(cmd *cobra.Command, args []string) {
		nType, _ := cmd.Flags().GetString("type")

		dockerComposeFile := ""
		switch nType {
		case "full-main":
			dockerComposeFile = "./single_node/docker-compose.fullnode.main.yml"
		case "full-nil":
			dockerComposeFile = "./single_node/docker-compose.fullnode.nile.yml"
		case "witness-private":
			dockerComposeFile = "./single_node/docker-compose.witness.private.yml"
		default:
			fmt.Println("Error: type not supported", nType)
		}
		if yes, isDir := utils.PathExists(dockerComposeFile); !yes || isDir {
			fmt.Println("Error: file not exists or not a file:", dockerComposeFile)
		}
		if msg, err := utils.StopDockerCompose(dockerComposeFile); err != nil {
			fmt.Println("Error: ", err)
			return
		} else {
			fmt.Println("Stop done: ", msg)
		}
	},
}

func init() {
	runSingleCmd.AddCommand(runSingleStopCmd)
	NodeCmd.AddCommand(runSingleCmd)

	runSingleCmd.Flags().StringP(
		"type", "t", "",
		"Node type you want to deploy (required, available: full-main, full-nile, witness-private)")

	if err := runSingleCmd.MarkFlagRequired("type"); err != nil {
		log.Fatalf("Error marking type flag as required: %v", err)
	}
}
