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
func findDirSize(dir Dir) (size int) {
	for _, v := range dir.dirs {
		size += findDirSize(v)
	}
	for _, v := range dir.files {
		size += v.size
	}
	return
}
func findMin(dir Dir, neededSpace int, min *int) (size int) {
	for _, v := range dir.dirs {
		size += findMin(v, neededSpace, min)
	}
	for _, v := range dir.files {
		size += v.size
	}
	if size >= neededSpace && size < *min {
		*min = size
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
	homeDirSize := findDirSize(homeDir)
	min := 700000000
	remanining := 70000000 - homeDirSize
	findMin(homeDir, 30000000-remanining, &min)
	return min
}

func main() {
	fmt.Println(day7())
}
