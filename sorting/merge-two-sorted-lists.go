main pkg

import "fmt"




func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	tail := result

	for (list1 != nil) && (list2 != nil) {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}
	if list1 == nil {
		tail.Next = list2
	} else {
		tail.Next = list1
	}
	return result.Next
}

func main(){
  mergeTwoLists()
}
