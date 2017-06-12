package radix

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestSort(t *testing.T) {
	data := [...]string{"", "Hello", "foo", "fo", "xb", "xa", "bar", "foo", "f00", "%*&^*&^&", "***"}
	sorted := data[0:]
	sort.Strings(sorted)

	a := data[0:]
	Sort(a)
	if !reflect.DeepEqual(a, sorted) {
		t.Errorf(" got %v", a)
		t.Errorf("want %v", sorted)
	}
}

func TestSort1k(t *testing.T) {
	data := make([]string, 1<<10)
	for i := range data {
		data[i] = strconv.Itoa(i ^ 0x2cc)
	}

	sorted := make([]string, len(data))
	copy(sorted, data)
	sort.Strings(sorted)

	Sort(data)
	if !reflect.DeepEqual(data, sorted) {
		t.Errorf(" got %v", data)
		t.Errorf("want %v", sorted)
	}
}

func TestSortBible(t *testing.T) {
	var data []string
	f, err := os.Open("res/bible.txt")
	if err != nil {
		log.Fatal(err)
	}
	for sc := bufio.NewScanner(f); sc.Scan(); {
		data = append(data, sc.Text())
	}

	sorted := make([]string, len(data))
	copy(sorted, data)
	sort.Strings(sorted)

	Sort(data)
	if !reflect.DeepEqual(data, sorted) {
		for i, s := range data {
			if s != sorted[i] {
				t.Errorf("%v  got: %v", i, s)
				t.Errorf("%v want: %v\n\n", i, sorted[i])
			}
		}
		//t.Errorf(" got %v", "data")
		//t.Errorf("want %v", "sorted")
	}
}

func BenchmarkSortMsdBible(b *testing.B) {
	b.StopTimer()
	var data []string
	f, err := os.Open("res/bible.txt")
	if err != nil {
		log.Fatal(err)
	}
	for sc := bufio.NewScanner(f); sc.Scan(); {
		data = append(data, sc.Text())
	}

	a := make([]string, len(data))
	for i := 0; i < b.N; i++ {
		copy(a, data)
		b.StartTimer()
		Sort(a)
		b.StopTimer()
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func BenchmarkSortStringsBible(b *testing.B) {
	b.StopTimer()
	var data []string
	f, err := os.Open("res/bible.txt")
	if err != nil {
		log.Fatal(err)
	}
	for sc := bufio.NewScanner(f); sc.Scan(); {
		data = append(data, sc.Text())
	}

	a := make([]string, len(data))
	for i := 0; i < b.N; i++ {
		copy(a, data)
		b.StartTimer()
		sort.Strings(a)
		b.StopTimer()
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func BenchmarkSortMsd1K(b *testing.B) {
	b.StopTimer()
	data := make([]string, 1<<10)
	for i := range data {
		data[i] = strconv.Itoa(i ^ 0x2cc)
	}

	a := make([]string, len(data))
	for i := 0; i < b.N; i++ {
		copy(a, data)
		b.StartTimer()
		Sort(a)
		b.StopTimer()
	}
}

func BenchmarkSortStrings1K(b *testing.B) {
	b.StopTimer()
	data := make([]string, 1<<10)
	for i := range data {
		data[i] = strconv.Itoa(i ^ 0x2cc)
	}

	a := make([]string, len(data))
	for i := 0; i < b.N; i++ {
		copy(a, data)
		b.StartTimer()
		sort.Strings(a)
		b.StopTimer()
	}
}
