/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	for pa := headA; pa != nil; pa = pa.Next {
		for pb := headB; pb != nil; pb = pb.Next {
			if pa == pb {
				return pa
			}
		}
	}
	return nil
}
