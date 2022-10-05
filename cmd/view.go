package cmd

import (
	"d7/utils"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View a PKGBUILD",
	Long: `View a PKGBUILD
d7 view <PKGNAME>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			utils.PrntRed("Provide a package!")
			os.Exit(1)
		} else if len(args) > 1 {
			utils.PrntRed("Too many arguments!")
			os.Exit(1)
		}

		dirName := "/tmp/d7/cloned/" + args[0]

		_, err := os.Stat(dirName)

		if err != nil {
			if os.IsNotExist(err) {
				cloneUrl := "https://aur.archlinux.org/" + args[0] + ".git"
				gitArgs := []string{"git", "clone", cloneUrl, dirName}

				utils.RunCmd(gitArgs, dirName, false, true)
			}
		} else {
			utils.PrntRed("Folder already exists, updating PKGBUILD")
			utils.RunCmd([]string{"git", "pull"}, dirName, true, false)
		}

		content, err := ioutil.ReadFile(dirName + "/" + "PKGBUILD")
		if err != nil {
			utils.PrntRed("Error while reading file")
			utils.RunCmd([]string{"rm", "-rf", args[0]}, "/tmp/d7/cloned", true, false)
			os.Exit(1)
		}

		fmt.Println()
		fmt.Println(string(content))

		utils.RunCmd([]string{"rm", "-rf", args[0]}, "/tmp/d7/cloned", true, false)
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
