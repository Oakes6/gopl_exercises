package counter

import (
	"bufio"
	"strings"
)

// counter for lines
type LineCounter struct {
	lines int
}

func (lc *LineCounter) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	itr := 1
	for {
		adv, _, _ := bufio.ScanLines(p, false)
		if adv == 0 {
			lc.lines = itr
			return itr, nil
		}
		p = p[adv:]
		itr++
	}
}

// counter for words
type WordCounter struct {
	words int
}

func (wc *WordCounter) Write(p []byte) (int, error) {
	itr := 1
	for {
		adv, _, _ := bufio.ScanWords(p, false)
		if adv == 0 {
			wc.words = itr
			return itr, nil
		}
		p = p[adv:]
		itr++
	}
}

type WordCounterr struct {
	Count int
}

func (wc *WordCounterr) Write(p []byte) (int, error) {
	numOfWords := 0
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		numOfWords += 1

	}

	wc.Count += numOfWords
	return len(p), nil
}
