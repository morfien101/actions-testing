package main

import (
	"flag"
	"fmt"

	"github.com/morfien101/actions-testing/chat"
)

var (
	version = "development"

	flagVersion = flag.Bool("v", false, "Show the version.")
)

func main() {
	flag.Parse()
	if *flagVersion {
		fmt.Println(version)
		return
	}
	output := chat.Say("%s", "Hello Github Actions")
	fmt.Println(output)
}
