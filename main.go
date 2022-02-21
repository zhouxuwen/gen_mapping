package main

import (
	"flag"
	"fmt"
)

var (
	PackageName    = ""
	InputFilePath  = ""
	OutputFilePath = ""
)

func init() {
	flag.StringVar(&PackageName, "p", "", "package name")
	flag.StringVar(&InputFilePath, "i", "", "input json file  path")
	flag.StringVar(&OutputFilePath, "o", "", "output go file path")
}

func main() {
	flag.Parse()
	err := GenMappingFile()
	if err != nil {
		panic(err)
	}
	fmt.Println("finish...")
}
