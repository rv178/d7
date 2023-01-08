package utils

import (
	"fmt"
	"os"
	"os/exec"
)

const red string = "\x1b[31m"
const green string = "\x1b[32m"
const blue string = "\x1b[34m"
const bold string = "\x1b[1m"
const reset string = "\x1b[0m"

func PrntRed(s string, nl bool) {
	if nl {
		fmt.Printf("%s%s>>%s %s%s%s\n", bold, red, reset, bold, s, reset)
	} else {
		fmt.Printf("%s%s>>%s %s%s%s", bold, red, reset, bold, s, reset)
	}
}

func PrntGreen(s string, nl bool) {
	if nl {
		fmt.Printf("%s%s>>%s %s%s%s\n", bold, green, reset, bold, s, reset)
	} else {
		fmt.Printf("%s%s>>%s %s%s%s", bold, green, reset, bold, s, reset)
	}
}

func PrntBlue(s string, nl bool) {
	if nl {
		fmt.Printf("%s%s>>%s %s%s%s\n", bold, green, reset, bold, s, reset)
	} else {
		fmt.Printf("%s%s>>%s %s%s%s", bold, green, reset, bold, s, reset)
	}
}

func RunCmd(cmd []string, dir string, setDir bool, hidden bool) {

	var cmdArgs []string

	for i := 1; i < len(cmd); i++ {
		cmdArgs = append(cmdArgs, cmd[i])
	}

	c := exec.Command(cmd[0], cmdArgs...)

	if !hidden {
		c.Stdout = os.Stdout
		c.Stdin = os.Stdin
		c.Stderr = os.Stderr
	}

	if setDir {
		c.Dir = dir
	}

	if err := c.Run(); err != nil {
		os.Exit(1)
	}
}
