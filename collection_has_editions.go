package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"time"
)

var (
	dir string
	level int
)

func main() {
	has()
	switch
}

func getlevel(root string, dir string) int {
	s := strings.Replace(dir, root, "", 1)
	l := strings.Split(s, "/")
	filtered := []string{}
	for _, v := range l {
		if len(v) > 0 {
			filtered = append(filtered, v)
		}
	}
	return len(filtered)
}

func walkDir(p string, i fs.DirEntry, err error) error {
	if err != nil {
		panic(err)
	}
	if i.IsDir() {
		if dir == "" {
			dir = p
		}
		l := getlevel(dir, p)
		if l > level {
			level = l
		}
	}
	return nil
}

func haseditions(dir string) int {
	start := time.Now()
	dir = ""
	level = 0

	dir, _ := filepath.Abs(dir)

	err := filepath.WalkDir(dir, walkDir)
	if err != nil {
		panic(err)
	}

	fmt.Println(dir, level)

	// time.Sleep(100 * time.Millisecond)
	fmt.Println(time.Since(start))
}
