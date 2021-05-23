package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	directory := flag.String("dir", "", "Enter a directory")

	flag.Parse()

	if *directory == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	printTree(*directory, "")

}

func printTree(dir string, indent string) error {

	var files []string

	fi, err := os.Stat(dir)
	if err != nil {
		return fmt.Errorf("could not stat %s: %v", dir, err)
	}

	fmt.Println(fi.Name())
	if !fi.IsDir() {
		return nil
	}
	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("could not read dir %s: %v", dir, err)
	}

	for _, fi := range fis {
		if fi.Name()[0] != '.' {
			files = append(files, fi.Name())
		}
	}

	for i, file := range files {
		add := "│  "
		if i == len(files)-1 {
			fmt.Printf(indent + "└──")
			add = "   "
		} else {
			fmt.Printf(indent + "├──")
		}

		if err := printTree(filepath.Join(dir, file), indent+add); err != nil {
			return err
		}
	}

	return nil
}
