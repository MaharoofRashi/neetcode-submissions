type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{queue: []int{}}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)

	n := len(this.queue)
	for i := 0; i < n-1; i++ {
		front := this.queue[0]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, front)
	}
}

func (this *MyStack) Pop() int {
	front := this.queue[0]
	this.queue = this.queue[1:]
	return front
}

func (this *MyStack) Top() int {
	return this.queue[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
