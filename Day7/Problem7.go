package main

import (
	"fmt"
	"github.com/srinchow/adventOfCode/utils"
	"os"
	"strconv"
	"strings"
)

const maxDiskSpace = uint64(70000000)
const updateSpaceNeeded = uint64(30000000)

type directory struct {
	name     string
	size     uint64
	children []*directory
	parent   *directory
}

func getDir(name string, parent *directory) *directory {
	return &directory{name: name, size: 0, parent: parent, children: []*directory{}}
}

// 1443806
func main() {
	file, err := os.Open("./Day7/input.txt")
	if err != nil {
		fmt.Println(fmt.Sprintf("Error opening file %v", err))
		return
	}
	defer utils.CloseFile(file)

	lines := utils.ParseFile(file)[0]
	dirStructure := getDir("", nil)

	buildDirectoryStructure(lines, dirStructure)

	for dirStructure.parent != nil {
		dirStructure = dirStructure.parent
	}

	dirStructure = dirStructure.children[0]

	res := uint64(0)
	solve1(dirStructure, &res)
	fmt.Println(res)

	freeSpaceRemaining := maxDiskSpace - dirStructure.size
	extraSpaceNeeded := updateSpaceNeeded - freeSpaceRemaining
	toDelete := dirStructure.size
	solve2(dirStructure, extraSpaceNeeded, &toDelete)
	fmt.Println(toDelete)
}

func buildDirectoryStructure(lines []string, dirStructure *directory) {
	for i := 0; i < len(lines); i++ {
		command := lines[i]

		// move to given directory
		if strings.HasPrefix(command, "$ cd") {
			_, dir := parseCommand(strings.Fields(command))
			if dir == ".." {
				dirStructure = dirStructure.parent
			} else {
				found := false
				// current Parent already has this child attached and just need to move the pointer to it
				for _, child := range dirStructure.children {
					if child.name == dir {
						found = true
						dirStructure = child
						break
					}
				}

				if !found {
					newDir := getDir(dir, dirStructure)
					dirStructure.children = append(dirStructure.children, newDir)
					dirStructure = newDir
				}

			}
		}

		// printing directory structure add the same structure to current Directory dirStructure
		if strings.HasPrefix(command, "$ ls") {
			for i++; i < len(lines); i++ {
				commandTemp := lines[i]

				// ls command print ends
				if strings.HasPrefix(commandTemp, "$") {
					i--
					break
				}
				size, name := parseCommand(strings.Fields(commandTemp))
				// files or directory in the current dir
				currNewDir := &directory{name: name, size: uint64(size), children: []*directory{}, parent: dirStructure}
				dirStructure.children = append(dirStructure.children, currNewDir)
				for size > 0 && currNewDir.parent != nil {
					currNewDir.parent.size += uint64(size)
					currNewDir = currNewDir.parent
				}

			}
		}
	}
}

func solve1(structure *directory, sum *uint64) {
	if len(structure.children) == 0 {
		return
	}

	if len(structure.children) > 0 && structure.size <= 100000 {
		*sum += structure.size
	}

	for _, child := range structure.children {
		solve1(child, sum)
	}

}

func solve2(structure *directory, spaceNeeded uint64, deleteDirectorySize *uint64) {
	// not a directory
	if len(structure.children) == 0 {
		return
	}

	if structure.size > spaceNeeded && structure.size < *deleteDirectorySize {
		*deleteDirectorySize = structure.size
	}

	for _, child := range structure.children {
		solve2(child, spaceNeeded, deleteDirectorySize)
	}

}

func parseCommand(fields []string) (int, string) {
	size, err := strconv.Atoi(fields[0])
	if err != nil {
		return 0, fields[len(fields)-1]
	}
	return size, fields[len(fields)-1]
}
