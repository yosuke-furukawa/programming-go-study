package sorting

import (
	"fmt"
	"os"
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
			Length("3m38s"),
		},
		{
			"Go",
			"Moby",
			"Moby",
			1992,
			Length("3m37s"),
		},
		{
			"Go Ahead",
			"Alicia Keys",
			"As I Am",
			2007,
			Length("4m36s"),
		},
		{
			"Ready 2 Go",
			"Martin Solveig",
			"Smash",
			2011,
			Length("4m24s"),
		},
	}
	sort.Sort(SortTable{tracks, "Album", "Artist"})
	printTracks(tracks, os.Stdout)
}

func TestCustomSort2(t *testing.T) {
	fmt.Println("--CustomSort(Length, Year)--")
	var tracks = []*Track{
		{
			"Go",
			"Delialah",
			"From the Roots up",
			2012,
			Length("3m38s"),
		},
		{
			"Go",
			"Moby",
			"Moby",
			1992,
			Length("3m38s"),
		},
		{
			"Go Ahead",
			"Alicia Keys",
			"As I Am",
			2007,
			Length("4m36s"),
		},
		{
			"Ready 2 Go",
			"Martin Solveig",
			"Smash",
			2011,
			Length("4m24s"),
		},
	}
	sort.Sort(SortTable{tracks, "Length", "Year"})
	printTracks(tracks, os.Stdout)
}

func TestStableSort(t *testing.T) {
	fmt.Println("--Stable--")
	var tracks = []*Track{
		{
			"Go",
			"Delialah",
			"From the Roots up",
			2012,
			Length("3m38s"),
		},
		{
			"Go Ahead",
			"Alicia Keys",
			"As I Am",
			2007,
			Length("4m36s"),
		},
		{
			"Go",
			"Moby",
			"Moby",
			1992,
			Length("3m38s"),
		},
		{
			"Ready 2 Go",
			"Martin Solveig",
			"Smash",
			2011,
			Length("4m24s"),
		},
	}
	sort.Stable(SortTable{tracks, "", ""})
	printTracks(tracks, os.Stdout)
}
