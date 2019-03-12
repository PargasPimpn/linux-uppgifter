package main

import (
        "bufio"
        "fmt"
        "os"
)

func main() {
        scanner := bufio.NewScanner(os.Stdin)
        fmt.Print("Ge input:")
        for scanner.Scan() {
                fmt.Println("hej", scanner.Text())

        }
}

