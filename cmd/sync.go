package cmd

import (
	"d7/utils"
	"github.com/mikkeloscar/gopkgbuild"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strings"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Update all packages",
	Long: `Update all packages

d7 sync
`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			utils.PrntRed("This command does not take arguments!")
			os.Exit(1)
		}

		items, _ := ioutil.ReadDir("/tmp/d7/cloned")

		for _, item := range items {
			if item.IsDir() {
				utils.PrntBlue("Updating " + item.Name())
				dirName := "/tmp/d7/cloned/" + item.Name()
				_, err := os.Stat(dirName)

				if err != nil {
					if os.IsNotExist(err) {
						cloneUrl := "https://aur.archlinux.org/" + item.Name() + ".git"
						gitArgs := []string{"git", "clone", cloneUrl, dirName}

						utils.PrntBlue("Cloning " + item.Name() + " to " + dirName)

						utils.RunCmd(gitArgs, dirName, false, false)

					}
				} else {
					utils.PrntRed("Folder already exists, running git pull")
					utils.RunCmd([]string{"git", "pull"}, dirName, true, false)
				}

				pkgb, err := pkgbuild.ParseSRCINFO(dirName + "/.SRCINFO")
				if err != nil {
					utils.PrntRed("Error parsing .SRCINFO")
					os.Exit(1)
				}

				var makeDeps []string
				for _, buildDep := range pkgb.Makedepends {
					makeDeps = append(makeDeps, buildDep.Name)
				}

				if len(makeDeps) != 0 {
					utils.PrntBlue("Build dependencies: " + strings.Join(makeDeps, ", "))

					pacmanArgs := []string{"sudo", "pacman", "-S"}

					pacmanArgs = append(pacmanArgs, makeDeps...)
					utils.RunCmd(pacmanArgs, dirName, true, false)
				}

				utils.PrntGreen("Building " + args[0])
				utils.RunCmd([]string{"makepkg", "-sci"}, dirName, true, false)

				if len(makeDeps) != 0 {
					utils.PrntBlue("Cleaning up..")

					pacmanRnsArgs := []string{"sudo", "pacman", "-Rns"}
					pacmanRnsArgs = append(pacmanRnsArgs, makeDeps...)

					utils.RunCmd(pacmanRnsArgs, dirName, true, false)
				}

				utils.PrntGreen("Building " + item.Name())
				utils.RunCmd([]string{"makepkg", "-sci"}, dirName, true, false)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
