type TimeMap struct {
	store map[string][]Entry
}

type Entry struct {
	value string
	timestamp int
}

func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string][]Entry),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key] = append(this.store[key], Entry{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	entries := this.store[key]

	if len(entries) == 0 {
		return ""
	}

	result := ""
	low, high := 0, len(entries) - 1

	for low <= high{
		mid := (low + high) / 2

		if entries[mid].timestamp <= timestamp {
			result = entries[mid].value
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return result
}
