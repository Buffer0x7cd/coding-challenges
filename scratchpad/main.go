package main

import "fmt"

func main() {
    num := 3

    switch num {
    case 1, 2, 3:
        fmt.Println("One, Two, or Three")
    case 4, 5, 6:
        fmt.Println("Four, Five, or Six")
    default:
        fmt.Println("Default")
    }
}
