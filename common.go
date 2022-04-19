// Copyright 2020 Joshua J Baker. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
package btree

const (
	degree   = 128
	maxItems = degree*2 - 1 // max items per node. max children is +1
	minItems = maxItems / 2
)

type cow struct {
	_ int // cannot be an empty struct
}
