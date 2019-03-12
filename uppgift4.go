package main

import (
	"fmt"
	"os"
	"sort"
)

var argument = os.Args

func main() {
	sort.Strings(argument)
	for _, value := range argument {
		fmt.Println(value)
	}
}
