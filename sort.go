// Package radix contains a string sorting algorithm.
//
// A fast string sorting algorithm
//
// The radix.Sort function is an optimized radix sort
// equivalent to sort.Strings in the Go standard library.
//
package radix

// Sort sorts a slice of strings in increasing byte-wise lexicographic order.
// It's equivalent sort.Strings in the standard library.
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

type list struct {
	str  string
	next *list
}

// Type frame represents a stack frame.
type frame struct {
	head, tail *list
	size       int // size of list, 0 if already sorted
	pos        int // current position in string
}

type bucket struct {
	head, tail *list
	size       int // size of list, 0 if already sorted
}

// intoBucket0 puts a list of elements into a bucket.
func intoBucket0(b *bucket, head, tail *list, size int) {
	if b.head != nil {
		b.tail.next = head
		b.tail = tail
		b.size += size
		return
	}
	b.head = head
	b.tail = tail
	b.size = size
}

// intoBucketa puts a list of elements into a bucket.
// The minimum and maximum character seen so far (chMin, chMax)
// are updated when the bucket is updated for the first time.
func intoBucket1(b *bucket, head, tail *list, size int,
	ch int, chMin, chMax *int) {
	if b.head != nil {
		b.tail.next = head
		b.tail = tail
		b.size += size
		return
	}
	b.head = head
	b.tail = tail
	b.size = size
	if ch < *chMin {
		*chMin = ch
	}
	if ch > *chMax {
		*chMax = ch
	}
}

// ontoStack puts the list in a bucket onto the stack.
// If the list contains at most insertBreak elements, its sorted
// with insertion sort. If both the the list on top of the stack
// and the list to be added to the stack are already sorted,
// the new list is appended to the end of the list on the stack
// and no new stack record is created.
func ontoStack(stack []frame, b *bucket, pos int) []frame {
	b.tail.next = nil
	if b.size <= insertBreak {
		if b.size > 1 {
			b.head, b.tail = insertSort(b.head, pos)
		}
		b.size = 0 // Mark as sorted.
	}
	if top := len(stack) - 1; b.size == 0 && top >= 0 && stack[top].size == 0 {
		stack[top].tail.next = b.head
		stack[top].tail = b.tail
	} else {
		stack = append(stack, frame{
			head: b.head,
			tail: b.tail,
			size: b.size,
			pos:  pos,
		})
		b.size = 0
	}
	b.head = nil
	return stack
}

// Breakpoint for insertion sort.
const insertBreak = 20

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

// DEBUG
/*
func printBucket(b *bucket) {
	r := b.head
	for {
		fmt.Println(" ", r.str)
		if r == b.tail {
			break
		}
		r = r.next
	}
}
*/

// DEBUG
/*
func printStack(s []frame) {
	for i, f := range s {
		fmt.Println("frame", i)
		r := f.head
		for {
			fmt.Println(" ", r.str)
			if r == f.tail {
				break
			}
			r = r.next
		}
	}
}
*/
