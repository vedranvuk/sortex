// Copyright 2020 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package sortex

import (
	"fmt"
)

func ExampleSortedKeyNames() {
	var m = map[string]int{
		"e": 5,
		"d": 4,
		"c": 3,
		"b": 2,
		"a": 1,
	}
	for _, key := range SortedKeyNames(m) {
		fmt.Printf("%s: %d\n", key, m[key])
	}
	// Output: a: 1
	// b: 2
	// c: 3
	// d: 4
	// e: 5
}
