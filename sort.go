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

const (
	insertBreak = 20
	byteBreak   = 16000
)

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
			continue
		}
		stack = intoBuckets(stack, frame.head, frame.pos)
	}
	return res
}

// intoBuckets traverses a list and puts the elements into buckets
// according to the byte in position pos. The elements are moved in blocks
// consisting of strings that have a common byte in this position.
// We keep track of the minimum and maximum characters encountered.
// In this way we may avoid looking at some empty buckets when we traverse
// the buckets in order and push the lists onto the stack.
func intoBuckets(stack []frame, a *list, pos int) []frame {
	// 0 bytes
	var b0 bucket
	// 1 byte
	b1 := make([]bucket, 256)
	chMin, chMax := 255, 0

	t := a
	prevCh := -1
	if len(t.str) > pos {
		prevCh = int(t.str[pos])
	}
	size := 1
	for tn := t.next; tn != nil; t, tn = tn, tn.next {
		ch := -1
		if len(tn.str) > pos {
			ch = int(tn.str[pos])
		}
		size++
		if ch == prevCh {
			continue
		}
		if prevCh == -1 {
			intoBucket0(&b0, a, t)
		} else {
			intoBucket1(&b1[prevCh], a, t, size-1, prevCh, &chMin, &chMax)
		}
		a = tn
		prevCh = ch
		size = 1
	}
	if prevCh == -1 {
		intoBucket0(&b0, a, t)
	} else {
		intoBucket1(&b1[prevCh], a, t, size, prevCh, &chMin, &chMax)
	}

	if b0.head != nil {
		stack = ontoStack(stack, &b0, pos)
	}
	for i, max := int(chMin), int(chMax); i <= max; i++ {
		if b1[i].head != nil {
			stack = ontoStack(stack, &b1[i], pos+1)
		}
	}
	return stack
}

/*
// adaptivRadixSort sorts a list r with n elements.
func adaptiveRadixSort(a *list, n int) *list {
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
			continue
		}
		if frame.size <= byteBreak {
			stack = intoBuckets(stack, frame.head, frame.pos)
		} else {
			stack = intoBuckets2(stack, frame.head, frame.pos)
		}
	}
	return res
}

// readBytes reads 0-2 bytes from pos p to the end of string in s.
// Output format: bit 17-16: number of bytes, bit 15-0: value.
func readBytes(s string, p int) int {
	switch len(s) - p {
	case 0:
		return 0
	case 1:
		return (1 << 16) | int(s[p])
	default:
		return (1 << 17) | (int(s[p]) << 8) | int(s[p+1])
	}
}

// intoBuckets2 is very similar to the one-byte version intoBuckets.
// The main difference is that we use the arrays used1 and used 2
// to reduce the number of empty buckets that are inspected.
func intoBuckets2(stack []frame, a *list, pos int) []frame {
	// 0 bytes
	var b0 bucket
	// 1 byte
	b1 := make([]bucket, 256)
	chMin, chMax := 255, 0
	// 2 bytes
	b2 := make([]bucket, 256*256)
	used1 := make([]bool, 256)
	used2 := make([]bool, 256)

	t := a
	prevKey := readBytes(t.str, pos)
	size := 1
	for tn := t.next; tn != nil; t, tn = tn, tn.next {
		key := readBytes(tn.str, pos)
		size++
		if key == prevKey {
			continue
		}
		switch prevKey >> 16 {
		case 0:
			intoBucket0(&b0, a, t)
		case 1:
			ch := prevKey & 0xFF
			intoBucket1(&b1[ch], a, t, size-1, ch, &chMin, &chMax)
		default:
			ch := prevKey & 0xFFFF
			intoBucket2(&b2[ch], a, t, size-1, ch, used1, used2)
		}
		a = tn
		prevKey = key
		size = 1
	}
	switch prevKey >> 16 {
	case 0:
		intoBucket0(&b0, a, t)
	case 1:
		ch := prevKey & 0xFF
		intoBucket1(&b1[ch], a, t, size, ch, &chMin, &chMax)
	default:
		ch := prevKey & 0xFFFF
		intoBucket2(&b2[ch], a, t, size, ch, used1, used2)
	}

	// 0 bytes
	if b0.head != nil {
		stack = ontoStack(stack, &b0, pos)
	}

	// 1-2 bytes
	var lowByte []int
	for ch := 0; ch < 256; ch++ {
		if used2[ch] {
			lowByte = append(lowByte, ch)
		}
	}
	for ch1 := 0; ch1 < 256; ch1++ {
		if b1[ch1].head != nil {
			stack = ontoStack(stack, &b1[ch1], pos+1)
		}
		if !used1[ch1] {
			continue
		}
		high := ch1 << 8
		for _, ch2 := range lowByte {
			ch := high | ch2
			if b2[ch].head != nil {
				stack = ontoStack(stack, &b2[ch], pos+2)
			}
		}
	}
	return stack
}
*/

// intoBucket0 puts a list of elements into a bucket.
func intoBucket0(b *bucket, head, tail *list) {
	if b.head != nil {
		b.tail.next = head
		b.tail = tail
		return
	}
	b.head = head
	b.tail = tail
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

/*
// intoBucket2 puts a list of elements into a bucket.
// For two-byte bucketing the bookkeeping is more elaborate.
// We use two bool slices, used1 and used2, to keep track of
// what bytes occur in the first and second position.
func intoBucket2(b *bucket, head, tail *list, size int,
	ch int, used1, used2 []bool) {
	if b.head != nil {
		b.tail.next = head
		b.tail = tail
		b.size += size
		return
	}
	b.head = head
	b.tail = tail
	b.size = size
	used1[ch>>8] = true
	used2[ch&0xFF] = true
}
*/

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
