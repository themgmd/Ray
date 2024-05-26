package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

// addCmd represents the generate command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "An add is command for creating module",
	Long: `Create is command for creating module or app. For example:

- ray add module User
	`,
	PreRun: preRun,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(`specify command with module or app subcommand. For example:
- ray add module User
		`)
	},
}
