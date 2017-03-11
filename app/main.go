package main

import (
    "fmt"
    "time"
)

func main() {
    var hello string = "Hello World"
    fmt.Println( hello )

    hi := "Hi World"
    fmt.Println( hi )

    var mine = true
    fmt.Println( "The world is mine! True or false? ", mine )

    yours := ! true
    fmt.Println( "The world is mine! True or false? ", yours )

    var MC, ML, XXX = 1, 9, 80
    fmt.Println( "Since when is the world yours?", MC, ML, XXX )

    const say string = "wat?!"
    fmt.Println( "A constant is immutable!", say )

    fmt.Println( "\nLook, ma! I can haz count!" )
    for i := 0; i < 3; i++ {
        fmt.Println( i )
    }

    fmt.Println( "\nLook, ma! I can haz…" )
    for i := 0; i < 3; i++ {
        fmt.Println( i )
        if i%2 == 0 {
            fmt.Println( "…a break" )
            break;
        }
    }

    fmt.Println( "\nLet's see what is true and what is false:" )
    if num := 2; num < 0 {
        fmt.Println( num, "is less " )
    } else {
        fmt.Println( num, "is more" )
    }
    fmt.Println( "…than zero" )

    spaces := "spaces"
    fmt.Print( "\nPrint: This has no", spaces, "in the string" )
    fmt.Println( "\nPrintln: This has", spaces, "in the string\n" )

    switch time.Now().Weekday() {
    case time.Sunday, time.Saturday:
        fmt.Println( "Weekend!" )
    default:
        fmt.Println( "Worktime… *sight*" )
    }

    fmt.Println( "Defining for repition:\n" )
    WhatAreYou := func(i interface{}) {
        fmt.Println( "The type is…" )
        switch t := i.(type) {
        case bool :
            fmt.Println( t, "A boolean!" )
        case int :
            fmt.Println( t, "Surprise, I am an integer" )
        default:
            fmt.Println( t, "No idea who I am, %T\n" )
        }
    }
    WhatAreYou( true )
    WhatAreYou( 1 )
    WhatAreYou( "true" )

    fmt.Println( "Let's try an array" )
    var loop [5]int
    fmt.Println( "The contents are:", loop )
    fmt.Println( "The length is:", len( loop ) )
    loop[2] = 99
    fmt.Println( "The contents now are:", loop )
    setloop := [5]string{"Another", "Loop", "With", "Preset", "Data"}
    fmt.Println( "A preset loop", setloop )
}
