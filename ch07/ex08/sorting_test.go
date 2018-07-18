package ex08

import (
	"fmt"
	"sort"
	"testing"
)

func TestCustomSort(t *testing.T) {
	fmt.Println("--CustomSort(Album, Artist)--")
	var tracks = []*Track{
		{
			"Go",
			"Delialah",
			"From the Roots up",
			2012,
			length("3m38s"),
		},
		{
			"Go",
			"Moby",
			"Moby",
			1992,
			length("3m37s"),
		},
		{
			"Go Ahead",
			"Alicia Keys",
			"As I Am",
			2007,
			length("4m36s"),
		},
		{
			"Ready 2 Go",
			"Martin Solveig",
			"Smash",
			2011,
			length("4m24s"),
		},
	}
	sort.Sort(customSort{tracks, "Album", "Artist"})
	printTracks(tracks)
}

func TestCustomSort2(t *testing.T) {
	fmt.Println("--CustomSort(Length, Year)--")
	var tracks = []*Track{
		{
			"Go",
			"Delialah",
			"From the Roots up",
			2012,
			length("3m38s"),
		},
		{
			"Go",
			"Moby",
			"Moby",
			1992,
			length("3m38s"),
		},
		{
			"Go Ahead",
			"Alicia Keys",
			"As I Am",
			2007,
			length("4m36s"),
		},
		{
			"Ready 2 Go",
			"Martin Solveig",
			"Smash",
			2011,
			length("4m24s"),
		},
	}
	sort.Sort(customSort{tracks, "Length", "Year"})
	printTracks(tracks)
}

func TestStableSort(t *testing.T) {
	fmt.Println("--Stable--")
	var tracks = []*Track{
		{
			"Go",
			"Delialah",
			"From the Roots up",
			2012,
			length("3m38s"),
		},
		{
			"Go Ahead",
			"Alicia Keys",
			"As I Am",
			2007,
			length("4m36s"),
		},
		{
			"Go",
			"Moby",
			"Moby",
			1992,
			length("3m38s"),
		},
		{
			"Ready 2 Go",
			"Martin Solveig",
			"Smash",
			2011,
			length("4m24s"),
		},
	}
	sort.Stable(customSort{tracks, "", ""})
	printTracks(tracks)
}
