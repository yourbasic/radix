# Your basic radix sort [![GoDoc](https://godoc.org/github.com/yourbasic/radix?status.svg)][godoc-radix]

### A fast string sorting algorithm

This is an optimized sorting algorithm equivalent to `sort.Strings`
in the Go standard library. For string sorting, a carefully implemented
radix sort can be considerably faster than Quicksort, sometimes
**more than twice as fast**.

### MSD radix sort

![Radix sort](res/radix.png)

A discussion of **MSD radix sort**, its implementation and a comparison
with other well-known sorting algorithms can be found in
[Implementing radixsort][implradix]. In summary, MSD radix sort
uses O(n) extra space and runs in O(n+B) worst-case time,
where n is the number of strings to be sorted and B
is the number of bytes that must be inspected to sort the strings.

### Installation

Once you have [installed Go][golang-install], run the `go get` command
to install the `radix` package:

    go get github.com/yourbasic/radix


### Documentation

There is an online reference for the package at
[godoc.org/github.com/yourbasic/radix][godoc-radix].


### Roadmap

* The API of this library is frozen.
* Version numbers adhere to [semantic versioning][sv].

Stefan Nilsson – [korthaj](https://github.com/korthaj)

[godoc-radix]: https://godoc.org/github.com/yourbasic/radix
[golang-install]: http://golang.org/doc/install.html
[implradix]: https://www.nada.kth.se/~snilsson/publications/Radixsort-implementation/
[sv]: http://semver.org/
