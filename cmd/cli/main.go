package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	instanceName string
	configPath   string
)

var registerInstanceCmd = &cobra.Command{
	Use:   "register",
	Short: "Registers instance in the environment. Provide instance name and path to the configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			instanceName = cmd.Flag("instance").Value.String()
			configPath   = cmd.Flag("path").Value.String()
		)
		if instanceName == "" {
			log.Fatal("instance name should not be empty")
		}
		if configPath == "" {
			log.Fatal("path should not be empty")
		}
	},
}

var rootCmd = &cobra.Command{Use: "recommender"}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(registerInstanceCmd)
	registerInstanceCmd.Flags().StringVarP(&instanceName, "name", "n", "", "Name of the instance (required)")
	registerInstanceCmd.Flags().StringVarP(&instanceName, "path", "p", "", "Path to the PostgreSQL config file (required)")
	err := registerInstanceCmd.MarkFlagRequired("name")
	if err != nil {
		log.Fatal("registerInstanceCmd.MarkFlagRequired: ", err)
	}
	err = registerInstanceCmd.MarkFlagRequired("path")
	if err != nil {
		log.Fatal("registerInstanceCmd.MarkFlagRequired: ", err)
	}
}

func main() {
	Execute()
}
