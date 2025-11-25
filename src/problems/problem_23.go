package problems

func Problem_23() {
	lists := []*ListNode{
		GenerateLinkedList([]int{2, 3, 4, 5}),
	}
	PrintLinkedList(mergeKLists(lists))
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	ans := lists

	for len(ans) > 1 {
		ans = mergeProcess(ans)
	}
	return ans[0]
}

func mergeProcess(lists []*ListNode) []*ListNode {
	res := []*ListNode{}
	n := len(lists)
	i := 0
	if n%2 == 1 {
		res = append(res, lists[0])
		i = 1
	}
	for i <= len(lists)-1 {
		res = append(res, mergeTwoListsEff(lists[i], lists[i+1]))
		i += 2
	}
	return res
}
