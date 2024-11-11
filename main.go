package main

import (
	"fmt"

	"github.com/Basu008/Payroll/server/config"
)

func main() {
	c := config.GetConfig("default")
	fmt.Println(c)
}
