package getsystem

import (
	"fmt"
	"os/exec"
)

func Getsysteminfo() string {
	cmd := exec.Command("lsb_release", "-d")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(stdout)
}