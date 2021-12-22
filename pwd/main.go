package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args

	pwd, err := os.Getwd()

	if err != nil {
		fmt.Println("Error:", err)
	}

	if len(arguments) == 1 || arguments[1] != "-P" {
		return
	}

	fileinfo, err := os.Lstat(pwd)

	if fileinfo.Mode()&os.ModeSymlink != 0 {
		realpath, err := filepath.EvalSymlinks(pwd)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(realpath)
	}
}
