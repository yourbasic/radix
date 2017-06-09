package radix

import (
	"reflect"
	"sort"
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
