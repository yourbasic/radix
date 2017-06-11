package radix

import "fmt"

// msdRadixSort sorts a list r with n elements.
func msdRadixSort(a *list, n int) *list {
	if n < 2 {
		return a
	}
	var res *list
	stack := []frame{{head: a, size: n}}
	for len(stack) > 0 {
		top := len(stack) - 1
		frame := stack[top]
		stack = stack[:top]
		if frame.size == 0 { // already sorted
			frame.tail.next = res
			res = frame.head
		} else {
			stack = intoBuckets(stack, frame.head, frame.pos)
		}
	}
	return res
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

// intoBuckets traverses a list and puts the elements into buckets
// according to the byte in position pos. The elements are moved in blocks
// consisting of strings that have a common byte in this position.
// We keep track of the minimum and maximum characters encountered.
// In this way we may avoid looking at some empty buckets when we traverse
// the buckets in order and push the lists onto the stack.
func intoBuckets(stack []frame, a *list, pos int) []frame {
	b := make([]bucket, 257) // b[256] holds strings with length equal to pos.
	chMin, chMax := 255, 0

	t := a
	prevCh := 256
	if len(t.str) > pos {
		prevCh = int(t.str[pos])
	}
	size := 1
	for tn := t.next; tn != nil; t, tn = tn, tn.next {
		ch := 256
		if len(tn.str) > pos {
			ch = int(tn.str[pos])
		}
		size++
		if ch == prevCh {
			continue
		}
		intoBucket(&b[prevCh], a, t, size-1, prevCh, &chMin, &chMax)
		a = tn
		prevCh = ch
		size = 1
	}
	intoBucket(&b[prevCh], a, t, size, prevCh, &chMin, &chMax)

	if b[256].head != nil {
		b[256].size = 0 // Mark as already sorted.
		stack = ontoStack(stack, &b[256], pos)
	}
	for i, max := int(chMin), int(chMax); i <= max; i++ {
		if b[i].head != nil {
			stack = ontoStack(stack, &b[i], pos+1)
		}
	}
	return stack
}

// intoBucket puts a list of elements into a bucket.
// The minimum and maximum character seen so far (chMin, chMax)
// are updated when the bucket is updated for the first time.
func intoBucket(b *bucket, head, tail *list, size int,
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
	if ch == 256 {
		return
	}
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

// DEBUG
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

// DEBUG
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
