/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
	var vals []int
	for node := head; node != nil; node = node.Next {
		vals = append(vals, node.Val)
	}

	left, right := 0, len(vals)-1

	for left < right {
		if vals[left] != vals[right] {
			return false
		}
		left++
		right--
	}
	return true
}
