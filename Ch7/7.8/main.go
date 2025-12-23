// This file sorts data into multiple table orders and outputs the data to stdout
package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

type trackList struct {
	tracks []*Track
	arr    []string
}

func (tl *trackList) Len() int      { return len(tl.tracks) }
func (tl *trackList) Swap(i, j int) { tl.tracks[i], tl.tracks[j] = tl.tracks[j], tl.tracks[i] }
func (tl *trackList) Less(i, j int) bool {

	for _, key := range tl.arr {
		switch key {
		case "Title":
			if tl.tracks[i].Title != tl.tracks[j].Title {
				return tl.tracks[i].Title < tl.tracks[j].Title
			}
		case "Artist":
			if tl.tracks[i].Artist != tl.tracks[j].Artist {
				return tl.tracks[i].Artist < tl.tracks[j].Artist
			}
		case "Album":
			if tl.tracks[i].Album != tl.tracks[j].Album {
				return tl.tracks[i].Album < tl.tracks[j].Album
			}
		case "Year":
			if tl.tracks[i].Year != tl.tracks[j].Year {
				return tl.tracks[i].Year < tl.tracks[j].Year
			}
		case "Length":
			if tl.tracks[i].Length != tl.tracks[j].Length {
				return tl.tracks[i].Length < tl.tracks[j].Length
			}
		}
	}

	return false
}

func updateArray(tl sort.Interface, col string) {

	var newarr []string
	newarr = append(newarr, col)

	switch tl := tl.(type) {
	case *trackList:
		for _, c := range tl.arr {
			if c != col {
				newarr = append(newarr, c)
			}
		}
		tl.arr = newarr
	}
}

func NewTrackList(t []*Track) sort.Interface {
	tl := new(trackList)
	tl.tracks = t
	tl.arr = []string{"Title", "Artist", "Album", "Year", "Length"}
	return tl
}

func main() {
	fmt.Println("init tracklist")
	tl := NewTrackList(tracks)
	sort.Sort(tl)
	printTracks(tracks)

	fmt.Println("\nReverse(byArtist):")
	updateArray(tl, "Artist")
	sort.Sort(sort.Reverse(tl))
	printTracks(tracks)

	fmt.Println("\nbyYear:")
	updateArray(tl, "Year")
	sort.Sort(tl)
	printTracks(tracks)
}
