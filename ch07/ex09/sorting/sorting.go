package sorting

import (
	"fmt"
	"io"
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

func (t *Track) String() string {
	return fmt.Sprintf("{ title = %s, artist = %s, album = %s, year = %d, length = %s }", t.Title, t.Artist, t.Album, t.Year, t.Length)
}

func Length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track, w io.Writer) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(w, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

type SortTable struct {
	T         []*Track
	FirstKey  string
	SecondKey string
}

func (x SortTable) Len() int {
	return len(x.T)
}

func LessKey(x, y *Track, key string) bool {
	switch key {
	case "Title":
		return x.Title < y.Title
	case "Album":
		return x.Album < y.Album
	case "Artist":
		return x.Artist < y.Artist
	case "Year":
		return x.Year < y.Year
	case "Length":
		return x.Length < y.Length
	}
	return false
}

func (x SortTable) Less(i, j int) bool {
	return LessKey(x.T[i], x.T[j], x.FirstKey) ||
		!LessKey(x.T[j], x.T[i], x.FirstKey) && LessKey(x.T[i], x.T[j], x.SecondKey)
}

func (x SortTable) Swap(i, j int) {
	x.T[i], x.T[j] = x.T[j], x.T[i]
}

func (x *SortTable) Select(key string) {
	x.SecondKey, x.FirstKey = x.FirstKey, key
}
