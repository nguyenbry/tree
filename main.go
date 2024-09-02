package main

import (
	"fmt"
	"log"
	"os"
)

func depth(acc []string, path string, level int, lastChild bool) ([]string, error) {
	f, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	var prefix string

	// generate what to print before the name of the file
	for i := 0; i < level; i++ {
		isLastToken := i == level-1 // the last token beofre the name of the file/folder

		if isLastToken {
			if lastChild {
				// among my siblings, am I the last?
				prefix += "└── "

			} else {
				prefix += "├── "
			}

			continue
		}

		prefix += "│   "

	}

	// add new line to tree
	out := append(acc, prefix+f.Name())

	// base case
	if !f.IsDir() {
		return out, nil
	}

	entries, err := os.ReadDir(path)

	if err != nil {
		return nil, err
	}

	numChildren := len(entries)

	for i, entry := range entries {
		newOut, err := depth(out, path+"/"+entry.Name(), level+1, i == numChildren-1)

		if err != nil {
			return nil, err
		}

		// propagate
		out = newOut
	}

	return out, nil
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("no path provided")
	}

	lines, err := depth(make([]string, 0), args[0], 0, false) // doesnt matter if true or false

	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
