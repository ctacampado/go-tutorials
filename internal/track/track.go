package track

import (
	"time"
)

// Track struct
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

// ByArtist represet a playlist that will be sorted by artist name
type ByArtist []*Track

func (x ByArtist) Len() int           { return len(x) }
func (x ByArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x ByArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
