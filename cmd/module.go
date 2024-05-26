package cmd

import (
	"github.com/spf13/cobra"
	"github.com/themgmd/ray/internal/generator"
	"log"
)

func init() {
	addCmd.AddCommand(moduleCmd)
}

// moduleCmd represents the module command
var moduleCmd = &cobra.Command{
	Use:    "module",
	Short:  "module creates app modules",
	Long:   `create app modules for backend application`,
	PreRun: preRun,
	Run:    moduleExecute,
}

func moduleExecute(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatalln("must enter a name of module")
	}

	err := generator.GenerateModule(args[0])
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Module: %s created!\n", args[0])
}
