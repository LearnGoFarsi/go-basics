package main

import (
	"fmt"
	"sort"
	"time"
)

type Track struct {
	Singer string
	Name   string
	Year   int
	Length time.Duration
}

type byYear []*Track

func (b byYear) Len() int {
	return len(b)
}

func (b byYear) Less(i, j int) bool {
	return b[i].Year < b[j].Year
}

func (b byYear) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

type byLength []*Track

func (b byLength) Len() int {
	return len(b)
}

func (b byLength) Less(i, j int) bool {
	return int(b[i].Length) < int(b[j].Length)
}

func (b byLength) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {

	album := []*Track{
		{
			Singer: "Hichkas",
			Name:   "Dastasho mosht karde",
			Year:   1398,
			Length: 240,
		},

		{
			Singer: "M.R Shajarian",
			Name:   "Zabane Atash",
			Year:   1388,
			Length: 480,
		},

		{
			Singer: "Shahin Najafi",
			Name:   "Nagoftamat Naro",
			Year:   1381,
			Length: 260,
		},
	}

	sotableByLength := byLength(album)
	sort.Sort(sotableByLength)
	for _, t := range sotableByLength {
		fmt.Printf("Track name: %s, Length: %d \n", t.Name, t.Length)
	}

	fmt.Println("*****************")

	sotableByYear := byYear(album)
	sort.Sort(sotableByYear)
	for _, t := range sotableByYear {
		fmt.Printf("Track name: %s, Year: %d \n", t.Name, t.Year)
	}
}
