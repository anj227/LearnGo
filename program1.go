package main

import (
    "fmt"
    t "time"
    "math/rand"
)


// Functions
func multiplier(x int32, y int32) int32 {
  return x * y
}

func main() {
    // Single line comments.
    /*
        Block comments 
    */
    fmt.Println("Hello World")
    fmt.Println("From the land of Go!")
    fmt.Println("Time now is: " )
    fmt.Println(t.Now())

    // Types:
    // var or const 
    const earthsGravity = 9.8
    // basic data types: int, float, complex ( 7 + 2i)
    // numerical: bool (1), int8, int16, uint8, uint16.. (unsigned)
    // int goes from -128 to 127, uint8 goes from 0 to 256
    var myName string
    myName = "Somebody"
    fmt.Println("Hi, my name is: ", myName)
    // inferring data type:
    //     :=
    // nuclearMeltdownOccurring := true

    // Multiple variables:
    var part1, part2 string
    part1 = "To be..."
    part2 = "Not to be..."
    fmt.Println(part1, part2)

    // Define amount and unit below:
    amount, unit := 10, "doll hairs"
    fmt.Println(amount, unit, ", that's expensive...")

    var xyz string
    fmt.Println("-----")
    fmt.Println(xyz)
    fmt.Println("-----")

    // printf:
    animal1 := "cat"
    animal2 := "dog"
    fmt.Printf("Are you a %v or a %v person?\n", animal1, animal2)
    // %v (for strings), %T (prints type of variable), %d (for numbers), %f (float), 
    //  %.2f --> float with 2 decimal points 
    // fmt is the formatter package. 
    // We have other methods that don’t print strings, but format them instead 
    //   like fmt.Sprint() and fmt.Sprintln().

    // Get user response using scan:
    // var response string 
    // fmt.Scan(&response)

    // Conditions:
    // == != > < >=, <= 
    // && ==> and,  || (or),  ! (not)
    ready := true 
    if (ready) {
        fmt.Println("I am ready!")
    } else {
        fmt.Println("I am not ready...")
    }

    var target int = 10
    var guess int = 5
    if (target == guess){  // == != > < >=, <= 
        fmt.Println("Yes!")
    } else {
        fmt.Printf("No.. you are off by: %d. \n", (target - guess))
    }

    clothingChoice := "sweater"
    switch clothingChoice{
        case "shirt": 
            fmt.Println("You will feel cold!")
        case "jacket":
            fmt.Println("You will feel hot!")
        case "sweater":
            fmt.Println("That should be just fine!")
        default:
            fmt.Println("I give up!")
    }

    // Random # -- must have!
    fmt.Println(rand.Intn(100))
    // with seeding:
    rand.Seed(time.Now().UnixNano())
    fmt.Println(rand.Intn(100))

}