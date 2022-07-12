/*
sortByTarget keep 's' the same as 'target' in order
Only useful when you want to use like sort.Sort, because
as you already have 'target' you don't need this sort method
at all, just return the 'target'!
*/

package main

import (
	"reflect"
	"testing"
)

func sortByTarget[T any](s []*T, target []*T) {
	idxMap := make(map[*T]int)
	itemMap := make(map[int]*T)
	for idx, f := range s {
		idxMap[f] = idx
		itemMap[idx] = f
	}

	swapper := reflect.Swapper(s)
	for idx, item := range target {
		cidx := idxMap[item]
		idxFlow := itemMap[idx]

		swapper(idx, cidx)

		idxMap[item] = idx
		idxMap[idxFlow] = cidx
		itemMap[idx] = item
		itemMap[cidx] = idxFlow
	}
}

func Test_sortByTarget(t *testing.T) {
	type item struct {
		a int
	}

	items := []*item{{1}, {2}, {3}, {4}, {5}}
	target := []*item{items[2], items[3], items[1], items[4], items[0]}

	sortByTarget(items, target)
	reflect.DeepEqual(items, target) // true
}
