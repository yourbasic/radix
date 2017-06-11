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
			stack = oneByte(stack, frame.head, frame.pos)
		} else {
			stack = twoBytes(stack, frame.head, frame.pos)
		}
	}
	return res
}

// TODO
func oneByte(stack []frame, a *list, pos int) []frame {
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

// TODO
func twoBytes(stack []frame, a *list, pos int) []frame {
	return nil
}

// TODO.
func intoBucket2(b *bucket, head, tail *list, size int,
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
