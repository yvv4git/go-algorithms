package main

type Node struct {
	Val  string
	Len  int // Each node stores the distance from the starting node
	Next *Node
}

type Queue struct {
	in  *Node
	out *Node
}

func createQueue(start string) Queue {
	q := Queue{}
	q.Push(start, 0)
	return q
}

func (q *Queue) Push(val string, len int) {
	newNode := &Node{
		Val: val,
		Len: len,
	}
	if q.in == nil {
		q.in = newNode
		q.out = newNode
	} else {
		q.in.Next = newNode
		if q.in == q.out {
			q.out.Next = newNode
		}
		q.in = newNode
	}
}

func (q *Queue) Pop() (string, int) {
	val, l := q.out.Val, q.out.Len
	if q.out.Next == nil {
		q.out = nil
		q.in = nil
	} else {
		q.out = q.out.Next
	}
	return val, l
}

func (q *Queue) IsEmpty() bool {
	return q.out == nil
}

// Using this function, we check that there is
// no more than 1 mutation in the next gene.
func isNextMutation(curr, next string) bool {
	count := 0
	for k := 0; k < 8 && count <= 1; k++ {
		if curr[k] != next[k] {
			count++
		}
	}
	return count == 1
}

// Standard Breadth-First Search
func minMutationV3(startGene string, endGene string, bank []string) int {
	// We will put the visited mutations in the map
	visited := make(map[string]struct{})

	q := createQueue(startGene)

	for !q.IsEmpty() {
		val, l := q.Pop()
		visited[val] = struct{}{}

		for i := 0; i < len(bank); i++ {
			if _, ok := visited[bank[i]]; !ok && isNextMutation(val, bank[i]) {
				if bank[i] == endGene {
					return l + 1
				}
				q.Push(bank[i], l+1)
			}
		}
	}

	return -1
}
