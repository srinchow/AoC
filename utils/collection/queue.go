package collection

import "fmt"

type Queue struct {
	elements []int
}

func (this *Queue) Push(val int) {
	this.elements = append(this.elements, val)
}

func (this *Queue) pop() {
	if len(this.elements) == 0 {
		fmt.Println("Unable to pop from empty queue")
		return
	}
	this.elements = this.elements[1:]
}
