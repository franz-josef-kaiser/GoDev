/**
Examples of the Go language
 */
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
	arrayloop := [5]int{1, 2, 3, 4, 5}
	fmt.Println( "A preset loop", arrayloop )

	var multidem [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			multidem[ i ][ j ] = i + j
		}
	}
	fmt.Println( "Multidimensional Array", multidem )

	fmt.Println( "\nNow we take a look at … Slices. Which are just a subtype of arrays." )
	slice := make([]string, 3)
	fmt.Println( "We declared a slice of type 'Array' with a length of", len( slice ), slice )
	slice[0] = "0: First slice entry"
	slice[1] = "1: Second slice entry"
	slice[2] = "2: Third slice entry"
	fmt.Println( "The slice now has contents", slice )

	fmt.Println( "Appending to the slice, returns a new slice" )
	slice = append( slice, "3: A fourth entry" )
	fmt.Println( "The new slice has the following entries", slice )

	fmt.Println( "We just need a part of the slice, so we … slice it" )
	slicecopy := make( []string, len( slice ) )
	copy( slicecopy, slice )
	fmt.Println( "A copy of the slice:", slicecopy )
	fmt.Println( "A slice of the slice [2:3]", slice[2:3] )
	fmt.Println( "A slice of the slice [:3]", slice[:3] )
	fmt.Println( "A slice of the slice [2:]", slice[2:] )
	setloop := [5]string{"Another", "Loop", "With", "Preset", "Data"}
	fmt.Println( "A preset loop", setloop )
	setloop2 := [...]string{"This", "Loop", "counts", "itself"}
	fmt.Println( "A preset loop that counts itself", setloop2 )

	fmt.Println( "\nAnother thing here: Maps" )

	fmt.Println( "\nSlices and Ranges: Two different, but related things" )
	fmt.Println( "Talking about ranges…" )
	nums := [] int { 2, 3, 4 }
	summedup := 0
	for _, num := range nums {
		summedup += num
	}
	fmt.Println( "Summed up, we got:", summedup )
}
