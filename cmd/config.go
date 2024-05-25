/*
Copyright © 2024 Muhammad Gasanov <therealmgmd@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"ray/internal/config"
)

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().String("package-name", "", "set a package name")
}

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configurate your CLI app",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		flag := cmd.Flag("package-name")
		err := config.SetPackage(flag.Value.String())
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("package name saved in config")
	},
}
