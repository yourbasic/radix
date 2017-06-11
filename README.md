# Your basic radix sort

### Golang string sorting algorithm

This library contains an optimized string sorting algorithm that sorts
a slice of strings in increasing order.
It's equivalent to the `sort.Strings` function in the standard Go library,
but considerably faster.

### Adaptive radix sort

![Radix sort](res/radix.png)

The algorithm is implemented using **Adaptive radix sort**,
an optimized version of **MSD radix sort**.

A discussion of the algorithm, its implementation and a comparison with other
well-known sorting algorithms, both bit-based and comparison-based,
can be found in the paper [Implementing radixsort][implradix].
The paper concludes that, for string sorting, carefully implemented
radix sorting algorithms are considerably faster (often more than
twice as fast) than comparison-based methods, and on average
Adaptive radix sort was the fastest algorithm.


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
