/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	indicies := make(map[*ListNode]bool)
	curr := head

	for curr != nil {
		if _, exists := indicies[curr]; exists {
			return true
		}

		indicies[curr] = true

		curr = curr.Next
	}

	return false
    
}
