package ex08

import (
	"fmt"
	"os"
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
	tw.Flush()
}

type customSort struct {
	t         []*Track
	firstKey  string
	secondKey string
}

func (x customSort) Len() int {
	return len(x.t)
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

func (x customSort) Less(i, j int) bool {
	return LessKey(x.t[i], x.t[j], x.firstKey) ||
		!LessKey(x.t[j], x.t[i], x.firstKey) && LessKey(x.t[i], x.t[j], x.secondKey)
}

func (x customSort) Swap(i, j int) {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}

func (x customSort) Select(key string) {
	x.firstKey, x.secondKey = key, x.firstKey
}
