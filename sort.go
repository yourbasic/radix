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
	//r := msd(&mem[0], n)
	for i := range a {
		a[i] = r.str
		r = r.next
	}
}

// TODO SortByte sorts b in byte-wise lexicographic order.
/*
func SortByte(b []byte) {
	return
}
*/

const insertBreak = 16

type bucket struct {
	head, tail *list
	size       int // size of list, 0 if already sorted
}

type frame struct {
	head, tail *list
	size       int // size of list, 0 if already sorted
	pos        int // current position in string
}

type list struct {
	str  string
	next *list
}

// ontoStack puts the list in a bucket onto the stack.
// If the list contains at most insertBreak elements, its sorted
// with insertion sort. If both the the list on top of the stack
// and the list to be added to the stack are already sorted,
// the new list is appended to the end of the list on the stack
// and no new stack record is created.
func ontoStack(b *bucket, stack []frame, pos int) {
	b.tail.next = nil
	if b.size <= insertBreak {
		if b.size > 1 {
			b.head, b.tail = insertSort(b.head, pos)
		}
		b.size = 0 // Mark as sorted.
	}
	if b.size == 0 || len(stack) == 0 || stack[len(stack)-1].size == 0 {
		stack = append(stack, frame{
			head: b.head,
			tail: b.tail,
			size: b.size,
			pos:  pos,
		})
	} else {
		top := stack[len(stack)-1]
		top.tail.next = b.head
		top.tail = b.tail
	}
	b.head = nil
}

// intoBucket puts a list of elements into a bucket.
func intoBucket(b *bucket, head, tail *list, size int, ch byte) {
	if b.head == nil {
		b.head = head
		b.tail = tail
		b.size = size
	} else {
		b.tail.next = head
		b.tail = tail
		b.size += size
	}
}

// intoBuckets traverses a list and puts the elements into buckets
// according to the character in position pos. The elements are moved
// in blocks consisting of strings that have a common character
// in position pos. We keep track of the minimum and maximum nonzero
// characters encountered. In this way we may avoid looking at some
// empty buckets when we traverse the buckets in ascending order
// and push the lists onto the stack.
func intoBuckets(a *list, stack []frame, pos int) {
	b := make([]bucket, 256)
	chMin, chMax := byte(255), byte(0)
	size := 1
	t := a
	prevCh := t.str[pos]
	for tn := t.next; tn != nil; t, tn = tn, tn.next {
		ch := tn.str[pos]
		size++
		if ch == prevCh {
			continue
		}
		intoBucket(&b[prevCh], a, t, size-1, prevCh)
		if prevCh <= chMin {
			chMin = prevCh
		}
		if prevCh >= chMax {
			chMax = prevCh
		}
		a = tn
		prevCh = ch
		size = 1
	}
	intoBucket(&b[prevCh], a, t, size, prevCh)
	if prevCh < chMin {
		chMin = prevCh
	}
	if prevCh > chMax {
		chMax = prevCh
	}

	// TODO Fix end-of-string.
	if b[0].head != nil {
		if b[0].size == 0 { // Mark as already sorted.
			//ontoStack(b, stack, pos)
		}
	}
	for i := chMin; i <= chMax; i++ {
		if b[i].head != nil {
			//ontoStack(i, stack, pos+1)
		}
	}
}

func msd(r *list, n int) (res *list) {
	if n < 2 {
		return
	}
	var stack []frame
	stack = append(stack, frame{
		head: r,
		size: n,
	})
	for len(stack) > 0 {
		n := len(stack)
		s := stack[n-1]
		stack = stack[:n-1]
		if s.size == 0 { // Is already sorted.
			s.tail.next = res
			res = s.head
			continue
		}
		intoBuckets(s.head, stack, s.pos)
	}
	return
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
