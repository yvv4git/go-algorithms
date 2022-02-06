package linked_list

// Element - used as element of doubles linked list
type Element struct {
	Data     string
	Previous *Element
	Next     *Element
}

//LinkedList - used as linked list data structure
type LinkedList struct {
	Head *Element
	Tail *Element
}

// NewLinkedList - used for fill linked list from data array
func NewLinkedList(datas []string) *LinkedList {
	result := &LinkedList{}

	for _, data := range datas {
		result.Add(data)
	}

	return result
}

// Add - used for add new element to linked list: O(1).
func (l *LinkedList) Add(data string) {
	newElement := &Element{
		Data:     data,
		Previous: l.Tail,
	}

	if l.Head == nil {
		l.Head = newElement
	} else {
		l.Tail.Next = newElement
	}

	l.Tail = newElement
}

// Delete - used for delete element from linked list: O(1).
func (l *LinkedList) Delete(element *Element) {
	if element.Previous != nil {
		element.Previous.Next = element.Next
	}

	if element.Next != nil {
		element.Next.Previous = element.Previous
	}

	if l.Head == element {
		l.Head = element.Next
	}

	if l.Tail == element {
		l.Tail = element.Previous
	}
}

// Search - used for search element in linked list: O(n).
func (l *LinkedList) Search(want string) *Element {
	currentElement := l.Head

	for {
		if currentElement.Data == want {
			return currentElement
		}

		if currentElement == l.Tail {
			break
		}

		currentElement = currentElement.Next
	}

	return nil
}

// InsertAfter - used for insert element in linked list: O(1).
func (l *LinkedList) InsertAfter(insert, after *Element) {
	insert.Previous = after
	insert.Next = after.Next

	if after.Next != nil {
		after.Next.Previous = insert
	}
	after.Next = insert

	if l.Tail == after {
		l.Tail = insert
	}
}
