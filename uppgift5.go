package main

import (
	"fmt"
	"os"
	"sort"
)

var argument = os.Args[1:]

func main() {
	sort.Strings(argument)
	//var s []string = [1:]
	for _, value := range argument {
		fmt.Println(value)
	}
}
