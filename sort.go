// TODO: Package radix contains a string sorting algorithm.
//
package radix

// Sort sorts a in byte-wise lexicographic order.
func Sort(a []string) {
	n := len(a)
	if n < 2 {
		return
	}
	// Build a linked list.
	mem := make([]list, n)
	for i, s := range a {
		mem[i].str = s
		if i < n-1 {
			mem[i].next = &mem[i+1]
		}
	}

	res := msd(&mem[0], n)

	// Put elements back into slice.
	for i := range a {
		a[i] = res.str
		res = res.next
	}
}

// TODO SortByte sorts b in byte-wise lexicographic order.
/*
func SortByte(b []byte) {
	return
}
*/

const insertBreak = 16

type list struct {
	str  string
	next *list
}

// insertSort sorts r comparing strings starting at position p.
// It returns the head and the tail of the sorted list.
func insertSort(r *list, p int) (head, tail *list) {
	head, tail = r, r
	for r := r.next; r != nil; r = tail.next {
		s := r.str[p:]
		switch {
		case tail.str[p:] <= s: // Add to tail.
			tail = r
		case head.str[p:] >= s: // Add to head.
			tail.next = r.next
			r.next = head
			head = r
		default: // Insert into middle.
			t := head
			for t.next.str[p:] <= s {
				t = t.next
			}
			tail.next = r.next
			r.next = t.next
			t.next = r
		}
	}
	return
}
