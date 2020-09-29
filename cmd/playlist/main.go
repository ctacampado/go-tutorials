package main

import (
	"fmt"
	"go-tutorials/internal/track"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

func main() {
	var tracks = []*track.Track{
		{Title: "Go", Artist: "Delilah", Album: "From the Roots Up", Year: 2012, Length: length("3m38s")},
		{Title: "Go", Artist: "Moby", Album: "Moby", Year: 1992, Length: length("3m37s")},
		{Title: "Go Ahead", Artist: "Alicia Keys", Album: "Am I Am", Year: 2007, Length: length("4m36s")},
		{Title: "Ready 2 Go", Artist: "Martin Solveig", Album: "Smash", Year: 2011, Length: length("4m24s")},
	}
	printTracks(tracks)

	fmt.Println()

	sort.Sort(track.ByArtist(tracks))
	printTracks(tracks)
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

// Prints all tracks
func printTracks(tracks []*track.Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")

	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}
