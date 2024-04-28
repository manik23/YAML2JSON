/*
Copyright © 2024 manik.x.mahajan@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	defaultVERSION = "v1"
)

func getVersion() *cobra.Command {

	return &cobra.Command{
		Use:   "version",
		Short: "version of cli",
		Long:  "version of cli",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(defaultVERSION)
		},
	}

}

func generateCommand() *cobra.Command {

	rootCmd := &cobra.Command{
		Use:   "yamlToJson",
		Short: "cli help to convert YAML file to JSON",
		Long:  "cli help to convert YAML file to JSON",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Tool to convert YAML to JSON")
		},
	}
	rootCmd.AddCommand(getVersion())
	rootCmd.AddCommand(parse())
	return rootCmd
}

func Execute() {

	cmd := generateCommand()

	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
