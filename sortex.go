// Copyright 2020 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package sortex contains mostly reflect based additions to std sort package.
package sortex

import (
	"reflect"
	"sort"
)

// SortedKeyNames is a helper for sorted map iteration. Returns a sorted slice
// of key names of inmap which must be a map of strings to any type. If inmap
// is not such type or is nil an empty slice is returned.
//
// Example:
//  var m map[string]interface{}
//  for _, key := range SortedKeyNames(m) {
//		fmt.Printf("%v", m[key])
//  }
func SortedKeyNames(inmap interface{}) []string {
	var mapv = reflect.ValueOf(inmap)
	if !mapv.IsValid() {
		return nil
	}
	if mapv.Kind() != reflect.Map {
		return nil
	}
	var ret = make([]string, 0, mapv.Len())
	var iter = mapv.MapRange()
	for iter.Next() {
		if iter.Key().Kind() != reflect.String {
			return nil
		}
		ret = append(ret, iter.Key().String())
	}
	sort.Strings(ret)
	return ret
}
