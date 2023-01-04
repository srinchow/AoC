package collection

import "fmt"

type Queue struct {
	elements []int
}

func (q *Queue) Push(val int) {
	q.elements = append(q.elements, val)
}

func (q *Queue) Pop() int {
	if len(q.elements) == 0 {
		fmt.Println("Unable to pop from empty queue")
		return -1
	}
	ele := q.elements[0]
	q.elements = q.elements[1:]
	return ele
}

func (q *Queue) Length() int {
	return len(q.elements)
}
