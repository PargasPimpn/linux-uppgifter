package main

import (
	"fmt"
	"sort"
)

var städer = []string{"Åbo", "Tammerfors", "Helsingfors", "Pargas", "Rovaniemi", "Imatra", "Kuopio", "Hyvinge", "S:t Michel", "Uleåborg", "Vasa"}

func main() {
	sort.Strings(städer)
	for _, value := range städer {
		fmt.Println(value)

	}
}
