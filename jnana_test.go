package main

import "testing"

func BenchmarkSearchAllBookmarks100(b *testing.B) {
	// run the searchAllBookmarks function b.N times
	query := "emblica"
	for n := 0; n < b.N; n++ {
		searchAllBookmarks(query)
	}
}
