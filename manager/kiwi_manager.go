package main

import (
	"fmt"
	"kiwi/manager/config"
)

func main() {
	fmt.Println(config.GetConf())
}
