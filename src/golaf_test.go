package main

import (
	"math"
	"testing"
)

const gapPenalty = -2
const mismatchPenalty = -3
const matchScore = 3

func TestCMap(t *testing.T) {
	keys := []string{"reset", "bold", "underline", "strike", "italic", "red", "green", "yellow", "blue", "purple", "cyan", "white"}
	cMap := getColorMap()
	for _, key := range keys {
		_, ok := cMap[key]
		if !ok {
			t.Fatalf("Key '%s' not found in getColorMap", key)
		}
	}
}

func TestReverse(t *testing.T) {
	input := []rune{1, 2, 3}
	target := []rune{3, 2, 1}
	reverseRune(input)
	for i := 0; i < len(input); i++ {
		if input[i] != target[i] {
			t.Fatalf("Reverse function failed")
		}
	}
}

func TestFMatrixEmpty(t *testing.T) {
	pattern := ""
	target := "abcd"
	fm, mI, mJ := fMatrix(pattern, target, gapPenalty, mismatchPenalty, matchScore)
	if fMLen := len(fm); fMLen != 1 {
		t.Fatalf("fMatrix got length of [%d] instead of 1", fMLen)
	}
	fmFSTarget := (len(target) + 1)
	if fMFirstSlice := len(fm[0]); fMFirstSlice != fmFSTarget {
		t.Fatalf("Constructed FMatrix slice got length of [%d] instead of %d", fMFirstSlice, fmFSTarget)
	}
	if mI != 0 {
		t.Fatalf("maxI got [%d] instead of 0", mI)
	}
	if mJ != 0 {
		t.Fatalf("maxJ got [%d] instead of 0", mJ)
	}
	for cv, val := range fm[0] {
		if val != 0 {
			t.Fatalf("Value at position [%d] in FMatrix got [%d] instead of 0", cv, val)
		}
	}

	pattern, target = target, pattern
	fm, mI, mJ = fMatrix(pattern, target, gapPenalty, mismatchPenalty, matchScore)
	patterSize := len(pattern) + 1
	if fMLen := len(fm); fMLen != patterSize {
		t.Fatalf("fMatrix got length of [%d] instead of [%d]", fMLen, patterSize)
	}
	for cv, val := range fm {
		if valLen := len(val); valLen != 1 {
			t.Fatalf("Value of slice [%d] got length of [%d] instead of 1", cv, valLen)
		}
		if val[0] != 0 {
			t.Fatalf("Value in slice [%d] in FMatrix got [%d] instead of 0", cv, val)
		}
	}
	if mI != 0 {
		t.Fatalf("maxI got [%d] instead of 0", mI)
	}
	if mJ != 0 {
		t.Fatalf("maxJ got [%d] instead of 0", mJ)
	}
}

func TestFMatrixFull(t *testing.T) {
	pattern := "abd"
	target := "abcd"
	fm, mI, mJ := fMatrix(pattern, target, gapPenalty, mismatchPenalty, matchScore)
	fmWantLen := len(pattern) + 1
	if fmLen := len(fm); fmLen != fmWantLen {
		t.Fatalf("FMatrix length is [%d] instead of %d", fmLen, fmWantLen)
	}
	fmWant := [][]int{{0, 0, 0, 0, 0}, {0, 3, 1, 0, 0}, {0, 1, 6, 4, 2}, {0, 0, 4, 3, 7}}
	fmSliceWantLen := len(target) + 1
	for cv, val := range fmWant {
		if valLen := len(val); valLen != fmSliceWantLen {
			t.Fatalf("Slice %d got length of [%d] instead of %d", cv, valLen, fmSliceWantLen)
		}
	}
	for i := 0; i < fmWantLen; i++ {
		for j := 0; j < fmSliceWantLen; j++ {
			if fm[i][j] != fmWant[i][j] {
				t.Fatalf("Got [%d] at position (%d, %d) instead of %d", fm[i][j], i, j, fmWant[i][j])
			}
		}
	}
	if mI != 3 {
		t.Fatalf("maxI got [%d] instead of 3", mI)
	}
	if mJ != 4 {
		t.Fatalf("maxJ got [%d] instead of 4", mJ)
	}
}

func TestBacktraceEmpty(t *testing.T) {
	pattern := "abc"
	target := ""
	fm, mI, mJ := fMatrix(pattern, target, gapPenalty, mismatchPenalty, matchScore)
	a1, a2 := backtrace(fm, []rune(pattern), []rune(target), []rune{}, []rune{}, mI, mJ, gapPenalty, mismatchPenalty, matchScore)
	if a1Len := len(a1); a1Len != 0 {
		t.Fatalf("Alignemt1 got length of [%d] instead of 0", a1Len)
	}
	if a2Len := len(a2); a2Len != 0 {
		t.Fatalf("Alignemt2 got length of [%d] instead of 0", a2Len)
	}

	pattern, target = target, pattern
	fm, mI, mJ = fMatrix(pattern, target, gapPenalty, mismatchPenalty, matchScore)
	a1, a2 = backtrace(fm, []rune(pattern), []rune(target), []rune{}, []rune{}, mI, mJ, gapPenalty, mismatchPenalty, matchScore)
	if a1Len := len(a1); a1Len != 0 {
		t.Fatalf("Alignemt1 got length of [%d] instead of 0", a1Len)
	}
	if a2Len := len(a2); a2Len != 0 {
		t.Fatalf("Alignemt2 got length of [%d] instead of 0", a2Len)
	}
}

func TestBacktrace(t *testing.T) {
	pattern := "abd"
	target := "abcd"
	fm, mI, mJ := fMatrix(pattern, target, gapPenalty, mismatchPenalty, matchScore)
	a1, a2 := backtrace(fm, []rune(pattern), []rune(target), []rune{}, []rune{}, mI, mJ, gapPenalty, mismatchPenalty, matchScore)

	a1Want := []rune{100, 45, 98, 97}
	a2Want := []rune{100, 99, 98, 97}

	algnLenWant := int(math.Max(float64(len(pattern)), float64(len(target))))
	if a1Len := len(a1); a1Len != algnLenWant {
		t.Fatalf("Length of alignment is [%d] instead of %d", a1Len, algnLenWant)
	}
	if a2Len := len(a2); a2Len != algnLenWant {
		t.Fatalf("Length of alignment is [%d] instead of %d", a2Len, algnLenWant)
	}
	for cr, curRune := range a1 {
		if curRune != a1Want[cr] {
			t.Fatalf("Rune of alignment1 at index %d is [%d] instead of %d", cr, curRune, a1Want[cr])
		}
	}
	for cr, curRune := range a2 {
		if curRune != a2Want[cr] {
			t.Fatalf("Rune of alignment2 at index %d is [%d] instead of %d", cr, curRune, a2Want[cr])
		}
	}

}
