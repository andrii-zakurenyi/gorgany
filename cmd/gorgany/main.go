package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/andrii-zakurenyi/gorgany/pkg/version"
)

func main() {
	var versionFlag = flag.Bool("version", false, "")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("Gorgany %s\n", version.GitTag)
		return
	}

	fmt.Println("<- Gorgany ->")
	fmt.Printf("\t- gitTag:    %s\n", version.GitTag)
	fmt.Printf("\t- gitCommit: %s\n", version.GitCommit)
	fmt.Printf("\t- platform:  %s/%s\n", runtime.GOOS, runtime.GOARCH)
}
