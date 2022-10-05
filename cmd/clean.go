package cmd

import (
	"d7/utils"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean all build directories",
	Long: `Clean all build directories
d7 clean`,
	Run: func(cmd *cobra.Command, args []string) {

		items, _ := ioutil.ReadDir("/tmp/d7/cloned")
		for _, item := range items {
			utils.PrntRed("Removing " + item.Name())
			utils.RunCmd([]string{"rm", "-rf", item.Name()}, "/tmp/d7/cloned", true, false)
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
