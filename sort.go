// TODO: Package radix contains a string sorting algorithm.
//
package radix

// Sort sorts a slice of strings in increasing order.
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
	res := msdRadixSort(&mem[0], n)
	// Put elements back into slice.
	for i := range a {
		a[i] = res.str
		res = res.next
	}
}

// Breakpoint for insertion sort.
const insertBreak = 16

type list struct {
	str  string
	next *list
}

// insertSort sorts a list comparing strings starting at byte position p.
// It returns the head and the tail of the sorted list.
func insertSort(a *list, p int) (head, tail *list) {
	head, tail = a, a
	for r := a.next; r != nil; r = tail.next {
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
