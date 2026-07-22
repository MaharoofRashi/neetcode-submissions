type MyHashMap struct {
	data []int
}

func Constructor() MyHashMap {
    return MyHashMap{data: make([]int, 1000001)}
}

func (this *MyHashMap) Put(key int, value int) {
    this.data[key] = value + 1
}

func (this *MyHashMap) Get(key int) int {
    if this.data[key] == 0 {
		return -1
	}
	return this.data[key] - 1
}

func (this *MyHashMap) Remove(key int) {
    this.data[key] = 0
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */