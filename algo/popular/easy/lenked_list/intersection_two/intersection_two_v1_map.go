package intersection_two

func getIntersectionNodeV1(headA, headB *ListNode) *ListNode {
	/*
		METHOD: Using map.
		TIME COMPLEXITY: O(n + m)
		SPACE COMPLEXITY: O(n)
	*/
	if headA == nil || headB == nil {
		return nil
	}

	m := make(map[*ListNode]struct{})
	for {
		if headA != nil {
			if _, ok := m[headA]; ok {
				return headA
			}
			m[headA] = struct{}{}
			headA = headA.Next
		}
		if headB != nil {
			if _, ok := m[headB]; ok {
				return headB
			}
			m[headB] = struct{}{}
			headB = headB.Next
		}

		if headA == nil && headB == nil {
			break
		}
	}

	return nil
}
