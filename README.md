# Your basic radix sort

### A fast string sorting algorithm

This is an optimized sorting algorithm equivalent to `sort.Strings`
in the Go standard library. For string sorting, a carefully implemented
radix sort can be considerably faster than Quicksort, sometimes
**more than twice as fast**.

### MSD radix sort

![Radix sort](res/radix.png)

A discussion of **MSD radix sort**, its implementation and a comparison
with other well-known sorting algorithms can be found in
[Implementing radixsort][implradix].


### Installation

Once you have [installed Go][golang-install], run the `go get` command
to install the `radix` package:

    go get github.com/yourbasic/radix


### Roadmap

* The API of this library is not yet frozen.
* Version numbers will adhere to [semantic versioning][sv].

Stefan Nilsson – [korthaj](https://github.com/korthaj)

[godoc-radix]: https://godoc.org/github.com/yourbasic/radix
[golang-install]: http://golang.org/doc/install.html
[implradix]: https://www.nada.kth.se/~snilsson/publications/Radixsort-implementation/
[sv]: http://semver.org/
