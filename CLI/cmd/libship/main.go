package main

import (
	"fmt"

	"github.com/uvindusl/lib-ship/internal/getsystem"
)

func main() {
	fmt.Println("Welcome to lib-ship")

	systeminfo := getsystem.Getsysteminfo()

	fmt.Println(systeminfo)

}