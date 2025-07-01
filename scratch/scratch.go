package main

import (
	"bufio"
	"fmt"
	"os"
)

// TODO: finish frontmatter parsing.

func parseFrontmatter(p string) error {
	// read the file
	f, err := os.Open(p)
	if err != nil {
		return fmt.Errorf("Could not read the file: %w", err)
	}
	defer f.Close()
	// find the delimeters
	var l []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		l = append(l, s.Text())
	}

	fmt.Println(l)

	// extract and put into a struct? or list of bytes?

	return nil

}

func main() {
	p := "/Users/mason_cotter/dev/notes/test2.md"
	err := parseFrontmatter(p)
	if err != nil {
		fmt.Println("Error calling parseFrontmatter: ", err)
	}
}
