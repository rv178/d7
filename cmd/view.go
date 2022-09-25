package cmd

import (
	"d7/utils"
	"github.com/spf13/cobra"
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
		} else if len(args) > 1 {
			utils.PrntRed("Too many arguments!")
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

		utils.RunCmd([]string{"cat", "PKGBUILD"}, dirName, true, false)
		utils.RunCmd([]string{"rm", "-rf", args[0]}, "/tmp/d7/cloned", true, false)
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
