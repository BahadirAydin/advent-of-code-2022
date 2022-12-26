package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type File struct {
	name string
	size int
}

type Dir struct {
	name     string
	files    []File
	dirs     []Dir
	upperDir *Dir
}

func cd(homeDir *Dir, currentDir *Dir, path string) *Dir {

	if path == "/" {
		return homeDir
	} else if path == ".." {
		return currentDir.upperDir
	}
	for k, v := range currentDir.dirs {
		if v.name == path {
			return &(currentDir.dirs[k])
		}
	}
	return &Dir{
		name:  "",
		files: []File{},
		dirs:  []Dir{},
	}
}
func fillDir(currentDir *Dir, str string) {
	f := strings.Fields(str)
	if f[0] == "dir" {
		currentDir.dirs = append(currentDir.dirs, Dir{
			name:     f[1],
			files:    []File{},
			dirs:     []Dir{},
			upperDir: currentDir,
		})
	} else {
		fileSize, _ := strconv.Atoi(f[0])
		currentDir.files = append(currentDir.files, File{size: fileSize, name: f[1]})
	}
}
func findDirSize(dir Dir, list *[]int) (size int) {
	for _, v := range dir.dirs {
		size += findDirSize(v, list)
	}
	for _, v := range dir.files {
		size += v.size
	}
	if size <= 100000 {
		(*list) = append(*list, size)
	}
	return
}

func day7() int {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var homeDir Dir
	var currentDir *Dir
	currentDir = &homeDir
	for scanner.Scan() {
		command := scanner.Text()
		if command[0] == '$' && command[2] == 'c' {
			path := command[5:]
			currentDir = cd(&homeDir, currentDir, path)
		} else if command[0] == '$' && command[2] == 'l' {
			continue
		} else {
			fillDir(currentDir, command)
		}
	}
	list := make([]int, 0)
	findDirSize(homeDir, &list)
	sum := 0
	for _, v := range list {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(day7())
}
