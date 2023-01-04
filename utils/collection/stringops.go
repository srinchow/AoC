package collection

import (
	"fmt"
	"strconv"
)

func GetInt(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting string to integer")
		fmt.Println(err)
		return -1
	}
	return res
}
