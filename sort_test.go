package radix

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

var text = [...]string{"", "Hello", "foo", "fo", "bar", "foo", "f00", "%*&^*&^&", "***"}

var textSorted []string

func init() {
	data := text
	textSorted = data[0:]
	sort.Strings(textSorted)
}

func TestSort(t *testing.T) {
	data := text
	a := data[0:]
	Sort(a)
	if !reflect.DeepEqual(a, textSorted) {
		t.Errorf("sorted %v", textSorted)
		t.Errorf("   got %v", a)
	}
}

func TestSort1k(t *testing.T) {
	a := make([]string, 1<<10)
	for i := range a {
		a[i] = strconv.Itoa(i ^ 0x2cc)
	}
	Sort(a)
	if !sort.StringsAreSorted(a) {
		t.Errorf("got %v", a)
	}
}

func BenchmarkSortMSD1K(b *testing.B) {
	b.StopTimer()
	unsorted := make([]string, 1<<10)
	for i := range unsorted {
		unsorted[i] = strconv.Itoa(i ^ 0x2cc)
	}
	data := make([]string, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(data, unsorted)
		b.StartTimer()
		Sort(data)
		b.StopTimer()
	}
}

func BenchmarkSortStrings1K(b *testing.B) {
	b.StopTimer()
	unsorted := make([]string, 1<<10)
	for i := range unsorted {
		unsorted[i] = strconv.Itoa(i ^ 0x2cc)
	}
	data := make([]string, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(data, unsorted)
		b.StartTimer()
		sort.Strings(data)
		b.StopTimer()
	}
}
