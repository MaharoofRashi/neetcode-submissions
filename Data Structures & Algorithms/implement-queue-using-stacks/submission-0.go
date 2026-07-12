type MyQueue struct {
	queue []int
}

func Constructor() MyQueue {
	return MyQueue{queue: []int{}}
}

func (this *MyQueue) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyQueue) Pop() int {
	front := this.queue[0]
	this.queue = this.queue[1:]
	return front
}

func (this *MyQueue) Peek() int {
	return this.queue[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Peek();
 * param4 := obj.Empty();
 */
