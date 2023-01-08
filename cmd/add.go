package cmd

import (
	"d7/utils"
	"fmt"
	"os"
	"strings"

	"github.com/mikkeloscar/gopkgbuild"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an AUR package",
	Long: `Add an AUR package
d7 add <package_name>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			utils.PrntRed("Provide a package name!", true)
			os.Exit(1)
		} else if len(args) > 1 {
			utils.PrntRed("Too many arguments!", true)
			os.Exit(1)
		}

		dirName := "/tmp/d7/cloned/" + args[0]

		_, err := os.Stat(dirName)

		if err != nil {
			if os.IsNotExist(err) {
				cloneUrl := "https://aur.archlinux.org/" + args[0] + ".git"
				gitArgs := []string{"git", "clone", cloneUrl, dirName}

				utils.PrntBlue("Cloning "+args[0]+" to "+dirName, true)

				utils.RunCmd(gitArgs, dirName, false, false)

			}
		} else {
			utils.PrntRed("Folder already exists, running git pull", true)
			utils.RunCmd([]string{"git", "pull"}, dirName, true, false)
		}

		pkgb, err := pkgbuild.ParseSRCINFO(dirName + "/.SRCINFO")
		if err != nil {
			utils.PrntRed("Error parsing .SRCINFO", true)
			os.Exit(1)
		}

		var makeDeps []string
		for _, buildDep := range pkgb.Makedepends {
			makeDeps = append(makeDeps, buildDep.Name)
		}

		if len(makeDeps) != 0 {
			utils.PrntBlue("Build dependencies: "+strings.Join(makeDeps, ", "), true)

			utils.PrntBlue("Do you want to install them? [Y/n]: ", false)

			var response string

			_, err := fmt.Scanln(&response)

			if err != nil {
				os.Exit(1)
			}

			if response == "Y" || response == "y" {
				pacmanArgs := []string{"sudo", "pacman", "-S"}

				pacmanArgs = append(pacmanArgs, makeDeps...)
				utils.RunCmd(pacmanArgs, dirName, true, false)
			}
		}

		utils.PrntGreen("Building "+args[0], true)
		utils.RunCmd([]string{"makepkg", "-sci"}, dirName, true, false)

		if len(makeDeps) != 0 {
			utils.PrntBlue("Remove build dependencies? [Y/n]: ", false)

			var response string

			_, err := fmt.Scanln(&response)

			if err != nil {
				os.Exit(1)
			}

			if response == "Y" || response == "y" {
				pacmanRnsArgs := []string{"sudo", "pacman", "-Rns"}
				pacmanRnsArgs = append(pacmanRnsArgs, makeDeps...)
				utils.RunCmd(pacmanRnsArgs, dirName, true, false)
			} else {
				os.Exit(1)
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
