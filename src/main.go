package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
Get the color codes for terminal output

	:parameter
	:return
	*	colorMap: map that returns the terminal escape sequences to color strings
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
Reverse a slice containing runes in place

	:parameter
	*	inSlice: slice that should be reversed
	:return
*/
func reverseRune(inSlice []rune) {
	n := len(inSlice)
	for i := 0; i < n/2; i++ {
		inSlice[i], inSlice[n-1-i] = inSlice[n-1-i], inSlice[i]
	}
}

/*
Calculate the levenshtein distance matrix

	:parameter
	*	str1: pattern to search for
	*	str2: string to search in
	*	gapP: gap penalty
	*	mismatchP: mismatch penalty
	*	match: score for a match
	:return
	*	distMat: distance matrix between the strings
	*	maxI, maxJ: coordinates of the highest score in the distMat (axis 0, axis 1)
*/
func fMatrix(str1, str2 string, gapP, mismatchP, match int) ([][]int, int, int) {
	str1Slice := []rune(str1)
	str2Slice := []rune(str2)
	s1_len := len(str1Slice) + 1
	s2_len := len(str2Slice) + 1

	if s1_len == 0 {
		return [][]int{{len(str2)}}, 0, 0
	}
	if s2_len == 0 {
		return [][]int{{len(str1)}}, 0, 0
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
	return distMat, maxI, maxJ
}

/*
Trace back the path through the distance matrix recursively

	:parameters
	*	btDistMat: distance matrix to trace back the alignment
	*	s1: pattern to search for
	*	s2: string to search in
	*	algn1: storage slice for alignment of pattern to searchString
	*	algn2: storage slice for alignment of searchString to pattern
	*	i, j: coordinates of the highest score in the matrix
	*	gapP: gap penalty
	*	mismatchP: mismatch penalty
	*	match: score for a match
	:return
	* algn1, algn2: the filled alignments
*/
func backtrace(btDistMat [][]int, s1, s2, algn1, algn2 []rune, i, j, gapP, mismatchP, match int) ([]rune, []rune) {
	if btDistMat[i][j] > 0 {
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

	:parameter
	*	pattern: pattern to search for
	*	searchString: string to search in
	*	inAlgn1: alignment of pattern to searchString
	*	inAlgn2: alignment of searchString to pattern
	*	colorize: name of the color in getColorMap to highlight search results
	*	qualityCutOff: percentage of word needs to be found to count as a match
	:return
*/
func showSearch(pattern, searchString string, inAlgn1, inAlgn2 []rune, color string, qualityCutOff float32) string {
	cMap := getColorMap()
	_, ok := cMap[color]
	if !ok {
		log.Fatal("Invalif color choice:", color)
	}
	// number of gap runes in the original search pattern
	numGapRunePattern := float32(strings.Count(pattern, "-"))
	// length of the alignment
	lenMatch := float32(len(inAlgn2))
	// number of insertions into the pattern in the alignment
	numIns := float32(strings.Count(string(inAlgn1), "-"))
	// length of the not aligned pattern
	lenPattern := float32(len(pattern))
	// quality of the match
	quality := (lenMatch - (numIns - numGapRunePattern)) / (lenPattern - numGapRunePattern)

	if quality >= qualityCutOff {
		// search for aligned section in the search string and build regex pattern
		var rePatBuilder strings.Builder
		splitAlgn2 := strings.Split(string(inAlgn2), "-")
		partsNum := len(splitAlgn2)
		lastPartInd := partsNum-1
		for i := 0; i < partsNum; i++ {
			if len(splitAlgn2[i]) > 0 {
				// to avoid trailing *-?
				if i != lastPartInd {
					rePatBuilder.WriteString(splitAlgn2[i])
					rePatBuilder.WriteString("*-?")
				} else {
					rePatBuilder.WriteString(splitAlgn2[i])
				}
			}
		}
		m := regexp.MustCompile(rePatBuilder.String())
		allInd := m.FindAllSubmatchIndex([]byte(searchString), -1)
		// number of all matches
		numInds := len(allInd)
		// color all aligned sections
		var sb strings.Builder
		lastPrintInd := 0
		for ci, i := range allInd {
			if ci == 0 && i[0] > 0 {
				sb.WriteString(searchString[:i[0]])
			}
			sb.WriteString(cMap["bold"])
			sb.WriteString(cMap[color])
			sb.WriteString(searchString[i[0]:i[1]])
			sb.WriteString(cMap["reset"])
			// if there is a string between the current and the next match
			if ci < numInds-1 {
				if allInd[ci+1][0]-i[1] > 0 {
					sb.WriteString(searchString[i[1]:allInd[ci+1][0]])
				}
			}
			lastPrintInd = i[1]
		}
		sb.WriteString(searchString[lastPrintInd:])
		return sb.String()
	}
	return ""
}

/*
Parse command line arguments and execute search over files or stdin

	:parameters
	:return
*/
func argparse() {
	// optional arguments
	// gap penalty
	gapPenaltyPtr := flag.Int("gapp", -2, "gap penalty [NEGATIVE]")
	// mismatch penalty
	mmPenaltyPtr := flag.Int("mmp", -3, "missmatch penalty [NEGATIVE]")
	// match bonus
	matchBonusPtr := flag.Int("match", 3, "score for a match [POSITIVE]")
	// minimum required quality to count as a match
	qualityCutOffPtr := flag.Int("quality", 75, "percentage of the pattern that have to macht to be seen as match")
	// whether to color the output
	colorPtr := flag.String("color", "green", "color option for highlighting the found results- options: [ red green yellow blue purple cyan white ]")
	flag.Parse()
	quality := float32(*qualityCutOffPtr) / float32(100)
	// number of optional args
	numFalgsPassed := 0
	flag.Visit(func(f *flag.Flag) {
		numFalgsPassed++
	})
	numArgs := len(os.Args)
	if numArgs == 1 {
		fmt.Fprintf(os.Stderr, "No arguments supplied\n")
		os.Exit(1)
	}

	// where positional args start
	numArgSkip := (numFalgsPassed * 2) + 1
	// the pattern to search for
	searchPattern := os.Args[numArgSkip]
	// the files in which the search should be performed
	files := os.Args[numArgSkip+1:]
	// read from stdin
	if len(files) == 0 {
		buf := bufio.NewScanner(os.Stdin)
		lineCount := 0
		for buf.Scan() {
			line := buf.Text()
			fm, mI, mJ := fMatrix(searchPattern, line, *gapPenaltyPtr, *mmPenaltyPtr, *matchBonusPtr)
			ag1, ag2 := backtrace(fm, []rune(searchPattern), []rune(line), []rune{}, []rune{}, mI, mJ, *gapPenaltyPtr, *mmPenaltyPtr, *matchBonusPtr)
			reverseRune(ag1)
			reverseRune(ag2)

			searchStringRes := showSearch(searchPattern, line, ag1, ag2, *colorPtr, float32(quality))
			if len(searchStringRes) > 0 {
				fmt.Printf("%d: %s\n", lineCount, searchStringRes)
			}
			lineCount++
		}
		// read from file(s)
	} else {
		// colormap for terminal output
		cMap := getColorMap()
		for _, filepath := range files {
			fmt.Printf("%s%s%s%s%s\n", cMap["bold"], cMap["italic"], cMap["red"], filepath, cMap["reset"])
			file, err := os.Open(filepath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Can't open file %s\n", filepath)
				os.Exit(1)
			}
			defer file.Close()

			buf := bufio.NewScanner(file)
			lineCount := 0
			for {
				if !buf.Scan() {
					break
				}
				line := buf.Text()
				fm, mI, mJ := fMatrix(searchPattern, line, *gapPenaltyPtr, *mmPenaltyPtr, *matchBonusPtr)
				ag1, ag2 := backtrace(fm, []rune(searchPattern), []rune(line), []rune{}, []rune{}, mI, mJ, *gapPenaltyPtr, *mmPenaltyPtr, *matchBonusPtr)
				reverseRune(ag1)
				reverseRune(ag2)
				searchStringRes := showSearch(searchPattern, line, ag1, ag2, *colorPtr, float32(quality))
				if len(searchStringRes) > 0 {
					fmt.Printf("%d: %s\n", lineCount, searchStringRes)
				}
				lineCount++
			}
		}
	}
}
func main() {
	argparse()
	/*
		// the pattern to search for
		pattern := "pakcage"
		// where to search
		target := "xxxpackage-main blabla package"
		// target = "TGTTACGG"
		// pattern = "GGTTGACTA"
		gapPenalty := -2
		fm, mI, mJ := fMatrix(pattern, target, gapPenalty, -3, 3)
		a1, a2 := backtrace(fm, []rune(pattern), []rune(target), []rune{}, []rune{}, mI, mJ, gapPenalty, -3, 3)
		reverseRune(a1)
		reverseRune(a2)
		showSearch(pattern, target, a1, a2, "green", 0.6)
	*/
}
