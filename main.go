package main

import (
	_ "bufio"
	_ "flag"
	"fmt"
	_ "os"
	// "strings"
)

/*
Get the color codes for terminal output
*/
func getColorMap() map[string]string {
	colorMap := make(map[string]string)
	colorMap["reset"] = "\033[0m"
	colorMap["bold"] = "\033[1m"
	colorMap["underline"] = "\033[4m"
	colorMap["strike"] = "\033[9m"
	colorMap["italic"] = "\033[3m"
	colorMap["red"] = "\033[31m"
	colorMap["green"] = "\033[32m"
	colorMap["yellow"] = "\033[33m"
	colorMap["blue"] = "\033[34m"
	colorMap["purple"] = "\033[35m"
	colorMap["cyan"] = "\033[36m"
	colorMap["white"] = "\033[37m"
	return colorMap
}

/*
Calculate the levenshtein distance matrix
*/
func fMatrix(str1, str2 string, gapP, mismatchP, match int) ([][]int, int, int, int) {
	s1_len := len(str1) + 1
	s2_len := len(str2) + 1

	if s1_len == 0 {
		return [][]int{{len(str2)}}, 0, 0, 0
	}
	if s2_len == 0 {
		return [][]int{{len(str1)}}, 0, 0, 0
	}
	// create and pre fill the distance matrix
	distMat := make([][]int, 0, s1_len)
	for i := 0; i < s1_len; i++ {
		distMat = append(distMat, make([]int, 0, s2_len)[:s2_len])
	}

	maxScore := 0
	maxI := 0
	maxJ := 0
	// fill distance matrix
	str1Slice := []rune(str1)
	str2Slice := []rune(str2)
	for i := 1; i < s1_len; i++ {
		for j := 1; j < s2_len; j++ {
			// only pay mismatchP if it's a mismatch
			substitutionCost := mismatchP
			if str1Slice[i-1] == str2Slice[j-1] {
				substitutionCost = match
			}
			// insertion
			ins := distMat[i-1][j] + gapP
			// deletion
			del := distMat[i][j-1] + gapP
			// substitution or same
			sub := distMat[i-1][j-1] + substitutionCost
			// find biggest of these values and populate distance matrix
			switch {
			case ins >= del && ins >= sub:
				distMat[i][j] = ins
			case del >= ins && del >= sub:
				distMat[i][j] = del
			default:
				distMat[i][j] = sub
			}
			if distMat[i][j] < 0 {
				distMat[i][j] = 0
			}
			if distMat[i][j] > maxScore {
				maxScore = distMat[i][j]
				maxI = i
				maxJ = j
			}
		}
	}
	return distMat, maxScore, maxI, maxJ
}

/*
Trace back the path thorough the distance matrix recursively
*/
func backtrace(btDistMat [][]int, s1, s2, algn1, algn2 []rune, i, j, gapP, mismatchP, match int) ([]rune, []rune) {
	if btDistMat[i][j] > 0{
		cost := mismatchP
		if s1[i-1] == s2[j-1] {
			cost = match
		}
		if i > 0 && j > 0 && (btDistMat[i][j] == btDistMat[i-1][j-1]+cost) {
			algn1 = append(algn1, s1[i-1])
			algn2 = append(algn2, s2[j-1])
			i--
			j--
		} else if i > 0 && (btDistMat[i][j] == btDistMat[i-1][j]+gapP) {
			algn1 = append(algn1, s1[i-1])
			algn2 = append(algn2, []rune("-")[0])
			i--
		} else if j > 0 && (btDistMat[i][j] == btDistMat[i][j-1]+gapP) {
			algn1 = append(algn1, []rune("-")[0])
			algn2 = append(algn2, s2[j-1])
			j--
		}
		return backtrace(btDistMat, s1, s2, algn1, algn2, i, j, gapP, mismatchP, match)
	} else {
		return algn1, algn2
	}
}

/*
Compare the search results an print them in colors depending on what was found
*/

func main() {
	// the pattern to search for
	pattern := "The quick brown fox jumps over the lazy dog"
	// where to search
	target := "A furry brown fox in the box"
	target = "TGTTACGG"
	pattern = "GGTTGACTA"
	gapPenalty := -2
	fm, mS, mI, mJ := fMatrix(pattern, target, gapPenalty, -3, 3)
	for _, i := range fm {
		fmt.Println(i)
	}
	a1, a2 := backtrace(fm, []rune(pattern), []rune(target), []rune{}, []rune{}, mI, mJ, gapPenalty,-3,3)
	fmt.Println(mS, mI, mJ)
	fmt.Println(string(a1))
	fmt.Println(string(a2))

}
