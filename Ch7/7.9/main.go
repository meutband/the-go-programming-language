// This file sorts data into multiple table orders and outputs data to HTML page
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"text/template"
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

var trackTable = template.Must(template.New("Track").Parse(`
<h1> Tracks </h1>
<table>
<tr style='text-align: left'>
	<th onclick="submitform('Title')">Title
        <form action="" name="Title" method="post">
            <input type="hidden" name="orderby" value="Title"/>
        </form>
    </th>
	<th onclick="submitform('Artist')">Artist
        <form action="" name="Artist" method="post">
            <input type="hidden" name="orderby" value="Artist"/>
        </form>
    </th>
	<th onclick="submitform('Album')">Album
        <form action="" name="Album" method="post">
            <input type="hidden" name="orderby" value="Album"/>
        </form>
    </th>
	<th onclick="submitform('Year')">Year
        <form action="" name="Year" method="post">
            <input type="hidden" name="orderby" value="Year"/>
        </form>
    </th>
	<th onclick="submitform('Length')">Length
        <form action="" name="Length" method="post">
            <input type="hidden" name="orderby" value="Length"/>
        </form>
    </th>
</tr>
{{range .Tracks}}
<tr>
	<td>{{.Title}}</td>
	<td>{{.Artist}}</td>
	<td>{{.Album}}</td>
	<td>{{.Year}}</td>
	<td>{{.Length}}</td>
</tr>
{{end}}
</table>

<script>
function submitform(formname) {
    document[formname].submit();
}
</script>
`))

func printTracks(w io.Writer, tl sort.Interface) {
	switch tl := tl.(type) {
	case *trackList:
		trackTable.Execute(w, tl)
	}
}

type trackList struct {
	Tracks []*Track
	arr    []string
}

func (tl *trackList) Len() int      { return len(tl.Tracks) }
func (tl *trackList) Swap(i, j int) { tl.Tracks[i], tl.Tracks[j] = tl.Tracks[j], tl.Tracks[i] }
func (tl *trackList) Less(i, j int) bool {

	for _, key := range tl.arr {
		switch key {
		case "Title":
			if tl.Tracks[i].Title != tl.Tracks[j].Title {
				return tl.Tracks[i].Title < tl.Tracks[j].Title
			}
		case "Artist":
			if tl.Tracks[i].Artist != tl.Tracks[j].Artist {
				return tl.Tracks[i].Artist < tl.Tracks[j].Artist
			}
		case "Album":
			if tl.Tracks[i].Album != tl.Tracks[j].Album {
				return tl.Tracks[i].Album < tl.Tracks[j].Album
			}
		case "Year":
			if tl.Tracks[i].Year != tl.Tracks[j].Year {
				return tl.Tracks[i].Year < tl.Tracks[j].Year
			}
		case "Length":
			if tl.Tracks[i].Length != tl.Tracks[j].Length {
				return tl.Tracks[i].Length < tl.Tracks[j].Length
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
	tl.Tracks = t
	tl.arr = []string{"Title", "Artist", "Album", "Year", "Length"}
	return tl
}

func main() {
	fmt.Println("init tracklist")
	tl := NewTrackList(tracks)
	sort.Sort(tl)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {

		}

		for key, val := range r.Form {
			if key == "orderby" {
				updateArray(tl, val[0])
			}
		}
		sort.Sort(tl)
		printTracks(w, tl)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
