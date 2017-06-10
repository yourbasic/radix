// TODO: Package radix contains a string sorting algorithm.
//
package radix

// Sort sorts a in byte-wise lexicographic order.
func Sort(a []string) {
	n := len(a)
	if n < 2 {
		return
	}
	mem := make([]list, n)
	for i, s := range a {
		mem[i].str = s
		if i < n-1 {
			mem[i].next = &mem[i+1]
		}
	}
	r, _ := insertSort(&mem[0], 0)
	for i := range a {
		a[i] = r.str
		r = r.next
	}
}

// SortByte sorts b in byte-wise lexicographic order.
func SortByte(b []byte) {
	return
}

type bucket struct {
	head, tail *list
	size       int // size of list, 0 if already sorted
}

type stack struct {
	head, tail *list
	size       int // size of list, 0 if already sorted
	pos        int // current position in string
}

type list struct {
	str  string
	next *list
}

// insertSort sorts r comparing strings starting at position p.
// It returns the head and the tail of the sorted list.
func insertSort(r *list, p int) (head, tail *list) {
	head, tail = r, r
	for r := r.next; r != nil; r = tail.next {
		switch {
		case r.str[p:] >= tail.str[p:]: // Add to tail.
			tail = r
		case r.str[p:] <= head.str[p:]: // Add to head.
			tail.next = r.next
			r.next = head
			head = r
		default: // Insert into middle.
			t := head
			for r.str[p:] >= t.next.str[p:] {
				t = t.next
			}
			tail.next = r.next
			r.next = t.next
			t.next = r
		}
	}
	return
}
