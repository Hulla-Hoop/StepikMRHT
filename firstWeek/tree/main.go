package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	if !printFiles {
		err := Recurs(out, path, 0, 0)
		if err != nil {
			return err
		}
	} else {
		err := RecursFS(out, path, 0, 0)
		if err != nil {
			return err
		}
	}

	return nil
}

func Recurs(out io.Writer, path string, right int, left int) error {
	dir, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	var dirSl []string
	for _, entry := range dir {
		if entry.IsDir() {
			dirSl = append(dirSl, entry.Name())
		}
	}
	for i, entry.Name() := range dirSl {
		if i == len(dirSl)-1 {
			if right == 0 && left == 0 {
				fmt.Fprint(out, "└───", entry.Name(), "\n")
			} else {

				s := strings.Repeat("│	", right) + strings.Repeat("	", left)
				fmt.Fprint(out, s, "└───", entry.Name(), "\n")

			}

			err = Recurs(out, path+string(os.PathSeparator)+entry.Name(), right, left+1)
			if err != nil {
				return err
			}
		} else {
			if right == 0 && left == 0 {
				fmt.Fprint(out, "├───", entry.Name(), "\n")
			} else {
				s := strings.Repeat("│	", right) + strings.Repeat("	", left)

				fmt.Fprint(out, s, "├───", entry.Name(), "\n")
			}

			err = Recurs(out, path+string(os.PathSeparator)+entry.Name(), right+1, left)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func RecursFS(out io.Writer, path string, right int, left int) error {
	dir, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for i, entry := range dir {
		if entry.IsDir() {

		} else {
			f, err := os.Open(path + string(os.PathSeparator) + entry.Name())
			if err != nil {
				return err
			}
			s, err := f.Stat()
			if err != nil {
				return err
			}

		}

	}

	return nil
}
