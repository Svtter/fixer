package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func isDirectory(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return false, errors.New("os.Stat error")
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		// do directory stuff
		return true, nil
	case mode.IsRegular():
		// do file stuff
		return false, nil
	}
	return false, nil
}

func main() {
	fmt.Println("Start fix file and folders permission ...")
	var walkPath string
	if len(os.Args) > 1 {
		walkPath = os.Args[1]
	} else {
		walkPath = "."
	}

	err := filepath.Walk(walkPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if path == "." {
				return nil
			}

			fmt.Printf("fix %s...\n", path)
			if isDir, _ := isDirectory(path); isDir {
				fmt.Println("not directory.")
				err = os.Chmod(path, 0755)
			} else {
				fmt.Println("directory.")
				err = os.Chmod(path, 0644)
			}
			if err != nil {
				panic(err)
			}
			fmt.Println(path, info.Size())
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Done.")
}
