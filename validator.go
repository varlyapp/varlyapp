package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

var (
	dir   string
	level int
)

type ValidationResponse struct {
	Level       int
	Valid       bool
	HasEditions bool
}

func validateCollection(dir string) ValidationResponse {
	level = 0

	dir, _ = filepath.Abs(dir)

	err := filepath.WalkDir(dir, dirwalker)
	if err != nil {
		panic(err)
	}

	return ValidationResponse{
		Level: level,
		Valid: level == 1 || level == 2,
		HasEditions: level == 2,
	}
}

func getlevel(root string, dir string) int {
	fmt.Println("Root ", root)
	fmt.Println("Dir ", dir)
	s := strings.Replace(dir, root, "", 1)
	fmt.Println("Clean Dir ", s)
	l := strings.Split(s, "/")
	filtered := []string{}
	for _, v := range l {
		if len(v) > 0 {
			filtered = append(filtered, v)
		}
	}
	return len(filtered)
}

func dirwalker(p string, i fs.DirEntry, err error) error {
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
