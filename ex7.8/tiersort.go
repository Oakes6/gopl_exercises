package tiersort

import "time"

type Song struct {
	Artist   string
	Title    string
	Duration time.Duration
}

type ByColumns struct {
	songs          []Song
	columns        []columnComparison
	maxComparisons int
}

func NewByColumns(songs []Song, max int) *ByColumns {
	return &ByColumns{songs, nil, max}
}

type columnComparison func(a, b *Song) comparison

type comparison int

const (
	eq comparison = iota
	lt
	gt
)

func (c *ByColumns) LessArtist(a, b *Song) comparison {
	if a.Artist > b.Artist {
		return gt
	} else if a.Artist < b.Artist {
		return lt
	} else {
		return eq
	}
}

func (c *ByColumns) LessTitle(a, b *Song) comparison {
	if a.Title > b.Title {
		return gt
	} else if a.Title < b.Title {
		return lt
	} else {
		return eq
	}
}

func (c *ByColumns) LessDuration(a, b *Song) comparison {
	if a.Duration > b.Duration {
		return gt
	} else if a.Duration < b.Duration {
		return lt
	} else {
		return eq
	}
}

func (c *ByColumns) Less(a, b int) bool {
	for _, f := range c.columns {
		cmp := f(&c.songs[a], &c.songs[b])
		switch {
		case cmp == eq:
			continue
		case cmp == lt:
			return true
		case cmp == gt:
			return false
		}
	}
	return false
}

func (c *ByColumns) Len() int {
	return len(c.songs)
}
func (c *ByColumns) Swap(a, b int) {
	c.songs[a], c.songs[b] = c.songs[b], c.songs[a]
}

func (c *ByColumns) Select(comp columnComparison) {
	c.columns = append([]columnComparison{comp}, c.columns...)

	if len(c.columns) > c.maxComparisons {
		c.columns = c.columns[:c.maxComparisons]
	}
}
