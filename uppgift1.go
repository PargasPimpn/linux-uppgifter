//uppg1//

package main

import (
	"fmt"
)

var städer = []string{"Åbo", "Tammerfors", "Helsingfors", "Pargas", "Rovaniemi", "Imatra", "Kuopio", "Hyvinge", "S:t Michel", "Uleåborg", "Vasa"}

func main() {
	for _, value := range städer {
		if value == "Åbo" {
		fmt.Println("luktar")
}
		fmt.Println(value)

	}
}
