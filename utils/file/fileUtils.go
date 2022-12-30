package file

import (
	"bufio"
	"fmt"
	"os"
)

func CloseFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Println(fmt.Sprintf("Error Closing file %v", err))
	}
}

func ParseFile(f *os.File) [][]string {
	sc := bufio.NewScanner(f)
	var arr [][]string
	var currArr []string
	for sc.Scan() {
		cal := sc.Text()
		if len(cal) == 0 {
			arr = append(arr, currArr)
			currArr = nil
		}
		currArr = append(currArr, cal)
	}
	if len(currArr) != 0 {
		arr = append(arr, currArr)
	}
	return arr
}
