package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GetFiles(path string, filter func(path string, name string) bool) []string {
	var files []string
	if stat, err := os.Stat(path); err == nil {
		if stat.IsDir() {
			if dir, err := ioutil.ReadDir(path); err == nil {
				for _, f := range dir {
					files = append(GetFiles(filepath.Join(path, f.Name()), filter), files...)
				}
			}
		} else {
			if filter(path, stat.Name()) {
				files = append(files, path)
			}
		}
	}
	return files
}

func ReadLines(file string) []string {
	return strings.Split(ReadString(file), "\n")
}

func ReadString(file string) string {
	if readFile, err := ioutil.ReadFile(file); err == nil {
		return string(readFile)
	} else {
		return ""
	}
}
