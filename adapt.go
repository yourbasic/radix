package radix

const byteBreak = 1500

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

// intoBuckets2 is very similar to the one byte version intoBuckets.
// The main difference is that we use the arrays used1 and used 2
// to reduce the number of empty buckets that are inspected.
func intoBuckets2(stack []frame, a *list, pos int) []frame {
	b := make([]bucket, 256*256) // TODO Fix end of string.
	used1 := make([]int, 256)
	used2 := make([]int, 256)

	t := a
	prevCh := 0 // TODO Read two bytes.
	//prevCh := 256
	//if len(t.str) > pos {
	//	prevCh = int(t.str[pos])
	//}
	size := 1
	for tn := t.next; tn != nil; t, tn = tn, tn.next {
		ch := 0 // TODO Read two bytes.
		//ch := 256
		//if len(tn.str) > pos {
		//	ch = int(tn.str[pos])
		//}
		size++
		if ch == prevCh {
			continue
		}
		intoBucket2(&b[prevCh], a, t, size-1, prevCh, used1, used2)
		a = tn
		prevCh = ch
		size = 1
	}
	intoBucket2(&b[prevCh], a, t, size, prevCh, used1, used2)

	var buckets1, buckets2 int
	for ch := 0; ch < 256; ch++ {
		if used1[ch] == 1 {
			used1[buckets1] = ch
			buckets1++
		}
		if used2[ch] == 1 {
			used2[buckets2] = ch
			buckets2++
		}
	}

	for ch1 := 0; ch1 < buckets1; ch1++ {
		high := used1[ch1] << 8
		for ch2 := 0; ch2 < buckets2; ch2++ {
			ch := high | used2[ch2]
			if b[ch].head != nil {
				if true == false { // TODO Check if end of string.
					b[ch].size = 0 // Mark as already sorted.
				}
				stack = ontoStack(stack, &b[ch], pos+2)
			}
		}
	}
	return stack
}

// intoBucket2 puts a list of elements into a bucket.
// For two-byte bucketing the bookkeeping is more elaborate.
// We use two integer arrays, used1 and used2, to keep track of
// what characters occur in the first and second position.
func intoBucket2(b *bucket, head, tail *list, size int,
	ch int, used1, used2 []int) {
	if b.head != nil {
		b.tail.next = head
		b.tail = tail
		b.size += size
		return
	}
	b.head = head
	b.tail = tail
	b.size = size
	used1[ch>>8] = 1
	used2[ch&255] = 1
}
