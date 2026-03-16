package getsystem

import (
	"fmt"
	"os/exec"
	"strings"
)

func Getsysteminfo() string {
	cmd := exec.Command("lsb_release", "-d")

	stdout, err := cmd.Output()
	
	systeminfo := string(stdout)
	
	systeminfo = strings.Trim(systeminfo, "Description:")
	
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(systeminfo)
}