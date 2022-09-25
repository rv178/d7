package cmd

import (
	"d7/utils"
	//"github.com/mikkeloscar/gopkgbuild"
	"github.com/spf13/cobra"
	"os"
	//"strings"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an AUR package",
	Long: `Add an AUR package
d7 add <package_name>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			utils.PrntRed("Provide a package name!")
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

				utils.PrntBlue("Cloning " + args[0] + " to " + dirName)

				utils.RunCmd(gitArgs, dirName, false, false)

			}
		} else {
			utils.PrntRed("Folder already exists, updating PKGBUILD")
			utils.RunCmd([]string{"git", "pull"}, dirName, true, false)
		}

		//utils.PrntBlue("Resolving dependencies..")

		//pkgb, err := pkgbuild.ParseSRCINFO(dirName + "/.SRCINFO")
		//if err != nil {
		//utils.PrntRed("Error parsing .SRCINFO")
		//os.Exit(1)
		//}

		//var makeDeps []string
		//for _, buildDep := range pkgb.Makedepends {
		//makeDeps = append(makeDeps, buildDep.Name)
		//}

		//utils.PrntBlue("Build dependencies: " + strings.Join(makeDeps, ", "))

		//pacmanArgs := []string{"sudo", "pacman", "-S", strings.Join(makeDeps, " ")}
		//utils.RunCmd(pacmanArgs, dirName, true, false)

		utils.PrntGreen("Building " + args[0])
		utils.RunCmd([]string{"makepkg", "-sci"}, dirName, true, false)

		//utils.PrntBlue("Cleaning up..")
		//pacmanRnsArgs := []string{"sudo", "pacman", "-Rns", strings.Join(makeDeps, " ")}
		//utils.RunCmd(pacmanRnsArgs, dirName, true, false)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
