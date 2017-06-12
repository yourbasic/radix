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
	// 0-1 byte
	b := make([]bucket, 257) // b[256] holds strings with 0 bytes.
	chMin, chMax := 255, 0

	// 2 bytes
	b2 := make([]bucket, 256*256)
	used1 := make([]int, 256)
	used2 := make([]int, 256)

	t := a
	prevCh := 0 // TODO Read two bytes.
	size := 1
	for tn := t.next; tn != nil; t, tn = tn, tn.next {
		ch := 0 // TODO Read two bytes.
		size++
		if ch == prevCh {
			continue
		}
		if true { // TODO Check #bytes.
			intoBucket(&b[prevCh], a, t, size-1, prevCh, &chMin, &chMax)
		} else {
			intoBucket2(&b2[prevCh], a, t, size-1, prevCh, used1, used2)
		}
		a = tn
		prevCh = ch
		size = 1
	}
	if true { // TODO Check #bytes.
		intoBucket(&b[prevCh], a, t, size-1, prevCh, &chMin, &chMax)
	} else {
		intoBucket2(&b2[prevCh], a, t, size-1, prevCh, used1, used2)
	}

	// 0 bytes
	if b[256].head != nil {
		b[256].size = 0 // Mark as already sorted.
		stack = ontoStack(stack, &b[256], pos)
	}

	// 1 byte
	for i, max := int(chMin), int(chMax); i <= max; i++ {
		if b[i].head != nil {
			stack = ontoStack(stack, &b[i], pos+1)
		}
	}

	// 2 bytes
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
			if b2[ch].head != nil {
				if true == false {
					b2[ch].size = 0 // Mark as already sorted.
				}
				stack = ontoStack(stack, &b2[ch], pos+2)
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
