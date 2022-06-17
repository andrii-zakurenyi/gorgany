package main

import (
	"flag"
	"fmt"
	"runtime"
)

var (
	gitTag    = "<unknown>"
	gitCommit = "<unknown>"
)

func main() {
	var versionFlag = flag.Bool("version", false, "")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("Gorgany %s\n", gitTag)
		return
	}

	fmt.Println("Gorgany")
	fmt.Printf("\t- gitTag:    %s\n", gitTag)
	fmt.Printf("\t- gitCommit: %s\n", gitCommit)
	fmt.Printf("\t- platform:  %s/%s\n", runtime.GOOS, runtime.GOARCH)
}
