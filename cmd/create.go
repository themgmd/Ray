/*
Copyright © 2024 Muhammad Gasanov <therealmgmd@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

// createCmd represents the generate command
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "Create is command for creating modules or app",
	Long: `Create is command for creating modules or app. For example:

- ray create app
- ray create module User
	`,
	PreRun: preRun,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(`specify command with module or app subcommand. For example:
- ray create app
- ray create module User
		`)
	},
}
