package cmd

import (
	"github.com/spf13/cobra"
	"github.com/themgmd/ray/internal/config"
	"log"
)

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().String("package-name", "", "set a package name")
}

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A config command configure your CLI app",
	Long:  `A config command represents configuration for your CLI app`,
	Run: func(cmd *cobra.Command, args []string) {
		flag := cmd.Flag("package-name")
		err := config.SetPackage(flag.Value.String())
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("package name saved in config")
	},
}
