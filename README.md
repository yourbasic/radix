# Your basic radix sort


### A fast string sorting algorithm

This is an optimized radix sort equivalent to `sort.Strings`
in the Go standard library.


### MSD radix sort

![Radix sort](res/radix.png)

The algorithm is implemented using an optimized version of **MSD radix sort**.

A discussion of the algorithm, its implementation and a comparison
with other well-known sorting algorithms can be found in the paper
[Implementing radixsort][implradix]. The paper concludes that,
for string sorting, carefully implemented radix sorting algorithms
are considerably faster, **often more than twice as fast**,
than comparison-based methods.

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
