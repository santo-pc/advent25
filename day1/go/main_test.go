package main

import (
	"testing"
)

func TestHandle(t *testing.T) {
	idx, crosses := Handle('L', 50, 4)
	assertHandle(t, 46, 0, idx, crosses)

	idx, crosses = Handle('R', 10, 5)
	assertHandle(t, 15, 0, idx, crosses)

	idx, crosses = Handle('L', 0, 5)
	assertHandle(t, 95, 1, idx, crosses)

	idx, crosses = Handle('R', 0, 99)
	assertHandle(t, 99, 0, idx, crosses)

	idx, crosses = Handle('R', 0, 100)
	assertHandle(t, 0, 1, idx, crosses)

	idx, crosses = Handle('L', 50, 150)
	assertHandle(t, 0, 2, idx, crosses)

	idx, crosses = Handle('L', 0, 1)
	assertHandle(t, 99, 1, idx, crosses)

	// 	idx, crosses = Handle('L', 50, 68)
	//	assertHandle(t, 82, 0, idx, crosses)
	//
	// // The dial is rotated L68 to point at 82; during this rotation, it points at 0 once.
	//
	//	idx, crosses = Handle('L', 30, 82)
	//	assertHandle(t, 52, 0, idx, crosses)
	//
	// The dial is rotated L30 to point at 52.
	// The dial is rotated R48 to point at 0.
	// The dial is rotated L5 to point at 95.
	// The dial is rotated R60 to point at 55; during this rotation, it points at 0 once.
	// The dial is rotated L55 to point at 0.
	// The dial is rotated L1 to point at 99.
	// The dial is rotated L99 to point at 0.
	// The dial is rotated R14 to point at 14.
	// The dial is rotated L82 to point at 32; during this rotation, it points at 0 once.
}

func assertHandle(
	t *testing.T,
	wantIndex int,
	wantCross int,
	gotIndex int,
	gotCross int,
) {
	t.Helper()

	if gotIndex != wantIndex || gotCross != wantCross {
		t.Fatalf(
			"got (index=%d, crosses=%d), want (index=%d, crosses=%d)",
			gotIndex, gotCross, wantIndex, wantCross,
		)
	}
}
