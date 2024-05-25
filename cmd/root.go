/*
Copyright © 2024 Muhammad Gasanov <therealmgmd@gmail.com>
*/

package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"ray/internal/config"
)

func init() {
	err := config.Init()
	if err != nil {
		log.Fatalln(err)
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:    "ray",
	Short:  "Ray is create-go-app CLI tool",
	Long:   `Ray is application template generator`,
	PreRun: preRun,
	Run:    func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func preRun(cmd *cobra.Command, args []string) {
	if config.Get().Package == "" {
		log.Fatalln("please set package name following command: ray config --package-name <name>")
	}
}
