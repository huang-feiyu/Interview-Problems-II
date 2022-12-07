package day2

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	// 1. hash map: old node -> new node
	if head == nil {
		return nil
	}
	old2new := make(map[*Node]*Node)
	curr := head
	for curr != nil {
		old2new[curr] = &Node{curr.Val, nil, nil}
		curr = curr.Next
	}

	// 2. copy next & random
	curr = head
	for curr != nil {
		newNode := old2new[curr]
		newNode.Next = old2new[curr.Next]
		newNode.Random = old2new[curr.Random]
		curr = curr.Next
	}

	return old2new[head]
}
