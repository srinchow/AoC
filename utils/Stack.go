package utils

type node struct {
	val  rune
	prev *node
}

type Stack struct {
	length int
	top    *node
}

func (this *Stack) Size() int {
	return this.length
}

func (this *Stack) Push(val rune) {
	newTop := &node{
		val:  val,
		prev: this.top,
	}
	this.top = newTop
	this.length++
}

func (this *Stack) Top() rune {
	return this.top.val
}

func (this *Stack) Pop() rune {
	if this.length == 0 {
		return rune(32)
	}
	val := this.top.val
	this.top = this.top.prev
	this.length--
	return val
}
