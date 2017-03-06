package main

import (
    "fmt"
)

func main() {
    var hello string = "Hello World"
    fmt.Println(hello)

    hi := "Hi World"
    fmt.Println(hi)

    var mine = true
    fmt.Println("The world is mine! True or false? ", mine)

    yours := ! true
    fmt.Println("The world is mine! True or false? ", yours)
}
