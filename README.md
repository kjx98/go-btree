# btree

[![GoDoc](https://godoc.org/github.com/kjx98/go-btree?status.svg)](https://godoc.org/github.com/kjx98/go-btree)

An efficient [B-tree](https://en.wikipedia.org/wiki/B-tree) implementation in Go, forked from tidwall(https://github.com/tidwall/btree), keep Map only.

## Features

- Support for Generics (Go 1.18+).
- `Map` types for ordered key-value maps,
- Fast bulk loading for pre-ordered data using the `Load()` method.
- `Copy()` method with copy-on-write support.
- Thread-safe operations.
- [Path hinting](PATH_HINT.md) optimization for operations with nearby keys.

## Using

To start using this package, install Go and run:

```sh
$ go get github.com/kjx98/go-btree
```

## B-tree types

This package includes the following types of B-trees:

- [`btree Map`](#btreemap):
A fast B-tree for storing ordered key value pairs.
Go 1.18+ 

### btree.Map

```go
// Basic
New()				// new K/V map w/ ordered key
MapNew()			// new K/V map w/ lessCmp func
Set(key, value)		// insert or replace an item
Get(key) (value, bool)   // get an existing item
Delete(key)			// delete an item
Len()				// return the number of items in the map

// Iteration
Scan(iter)         // scan items in ascending order
Reverse(iter)      // scan items in descending order
Ascend(key, iter)  // scan items in ascending order that are >= to key
Descend(key, iter) // scan items in descending order that are <= to key.
Iter()             // returns a read-only iterator for for-loops.

// Array-like operations
GetAt(index)       // returns the item at index
DeleteAt(index)    // deletes the item at index

// Bulk-loading
Load(key, value)   // load presorted items into tree
```

#### Example

```go
package main

import (
	"fmt"
	"github.com/kjx98/go-btree"
)

func main() {
	// create a map
	users := btree.New[string, string]()

	// add some users
	users.Set("user:4", "Andrea")
	users.Set("user:6", "Andy")
	users.Set("user:2", "Andy")
	users.Set("user:1", "Jane")
	users.Set("user:5", "Janet")
	users.Set("user:3", "Steve")

	// Iterate over the maps and print each user
	users.Scan(func(key, value string) bool {
		fmt.Printf("%s %s\n", key, value)
		return true
	})
	fmt.Printf("\n")

	// Delete a couple
	users.Delete("user:5")
	users.Delete("user:1")

	// print the map again
	users.Scan(func(key, value string) bool {
		fmt.Printf("%s %s\n", key, value)
		return true
	})
	fmt.Printf("\n")

	// Output:
	// user:1 Jane
	// user:2 Andy
	// user:3 Steve
	// user:4 Andrea
	// user:5 Janet
	// user:6 Andy
	//
	// user:2 Andy
	// user:3 Steve
	// user:4 Andrea
	// user:6 Andy
}
```

## Performance

This implementation was designed with performance in mind. 

[![asciicast](https://asciinema.org/a/489805.svg)](https://asciinema.org/a/489805)

```
goos: linux
goarch: amd64
pkg: github.com/kjx98/go-btree
cpu: Intel(R) Core(TM) i5-4200U CPU @ 1.60GHz
BenchmarkMapInsertSeq-4      	 6909127	       178.3 ns/op
BenchmarkMapInsertRandom-4   	 2277057	       535.8 ns/op
BenchmarkMapFind-4           	 2472787	       492.6 ns/op
BenchmarkMapDelete-4         	 2153862	       520.4 ns/op
BenchmarkMapDeleteLeft-4     	1000000000	         0.7902 ns/op
PASS
ok  	github.com/kjx98/go-btree	38.738s
```

## Contact

Jesse Kuang [@kjx98](http://mail.21cn.com/jkuang)<br/>
Josh Baker [@tidwall](http://twitter.com/tidwall)

## License

Source code is available under the MIT [License](/LICENSE).
