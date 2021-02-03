package main

import (
	"flag"
	"fmt"
)

var yamlConfigFile = flag.String("configFile", "", "Input your config path")

func main() {
	flag.Parse()
	fmt.Println(*yamlConfigFile)
}
