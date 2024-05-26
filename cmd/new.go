package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newCmd)
}

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "A new command creates go application with standard layer",
	Long:  `A new command creates go application with standard layer`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("must enter a name of app")
		}

		packageName := strings.ToLower(args[0])

		cloneCmd := fmt.Sprintf("git clone https://github.com/themgmd/gotemplate %s", packageName)
		execCmd := exec.Command("sh", "-c", cloneCmd)
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr

		err := execCmd.Run()
		if err != nil {
			log.Fatalln(err)
		}

		replaceCmd := fmt.Sprintf("sed -i '' 's/gotemplace/%s/g' *", packageName)
		replacePackageCmd := exec.Command("sh", "-c", replaceCmd)
		replacePackageCmd.Dir = fmt.Sprintf("./%s", packageName)
		replacePackageCmd.Stdout = os.Stdout
		replacePackageCmd.Stderr = os.Stderr

		err = replacePackageCmd.Run()
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf(
			"Application: %s created successfully\n\nPlease replace \"gotemplate\" in go.mod for your own name",
			packageName,
		)
	},
}
