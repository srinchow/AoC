package main

import "fmt"

func nSubsequentUniqueCharecters(input string, n int) int {

	if len(input) < n {
		return -1
	}

	currentState := make(map[byte]int)

	for i := 0; i < n; i++ {
		currentState[input[i]]++
	}

	if len(currentState) == n {
		return 0
	}

	startPtr := 0

	for i := n; i < len(input); i++ {
		currentState[input[startPtr]]--
		if currentState[input[startPtr]] == 0 {
			delete(currentState, input[startPtr])
		}
		startPtr++
		currentState[input[i]]++
		if len(currentState) == n {
			return i
		}
	}

	return -1
}

func main() {
	s := "abc"
	res := nSubsequentUniqueCharecters(s, 14)
	if res == -1 {
		fmt.Println("WTF")
	}
	fmt.Println(res)
}
