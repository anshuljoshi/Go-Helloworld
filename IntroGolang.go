// Every Go program is made up of packages.
// Programs start running in package main.
// Main refers to "executable". Others are libraries.
package main

// Declares libraries referenced in this source file
import (
    "fmt"       // The fmt package (shorthand for format)
                // implements formatting for input and output.
)

// main is special. it's the function that gets called program is executed.
func main() {
    // Println outputs a line to stdout.
    fmt.Println("Bello world!")

    // function call
    movingforward()
}

func movingforward() {
    var x int // Variable declaration.
    x = 10     // Variable assignment.
    y := 20    // "Short" declarations.
              // Go compiler is able to infer the type based on the literal
              // value assigned to the variable.
    sum, prod := returnTwoValues(x, y)        // Function returns two values.
    fmt.Println("sum:", sum, "\nprod:", prod) // Print.
    learnArrays()
}

/*  Functions can have parameters and multiple return values.
    Here `sum`, `prod` are the signature of what is returned. */
func returnTwoValues(x, y int) (sum, prod int) {
    return x + y, x * y // Return two values.
}

// Some built-in types and literals.
func learnArrays() {
    // Arrays have size fixed at compile time.
    var arr1 [5]int             // An array of 5 ints initialized to all 0.
    arr2 := [...]int{6,1,9}     // An array initialized with a fixed size of
                                // three elements

    fmt.Println(arr1)
    fmt.Println(arr2)
    // Slices have dynamic size.
    slc1 := []int{1,2,3}    // Compare to arr2. No ellipsis here.
    slc2 := make([]int, 5)    // Allocates slice of 5 ints, initialized to all 0.

    // Slices are dynamic.
    // To append elements to a slice, built-in append() function is used.
    slc1 = append(slc1, 4, 5, 6)   // Added 3 elements.
    fmt.Println(slc1)           // Updated slice is now [1 2 3 4 5 6]
    fmt.Println(slc2)
    // To append another slice to it.
    slc1 = append(slc1, []int{7, 8, 9}...) // Second argument is a slice literal.
    fmt.Println(slc1)           // Updated slice is now [1 2 3 4 5 6 7 8 9]

    p, q := learnPointers() // Declares p, q to be type pointer to int.
    fmt.Println(*p, *q)

    // Maps are a dynamically growable associative array type
    // (like the hash or dictionary types of Python or Java)
    m := map[string]int{"two": 2, "seven": 7}
    m["one"] = 1
    a:=10
    b:=20
    fmt.Println(learnNamedReturns(a,b))
    basicFlowControl() // Back in the flow.
}

// Assigning a name to the type being returned in the function declaration line
// allows to easily return from multiple points in a function
func learnNamedReturns(x, y int) (z int) {
    z = x * y
    return // z is implicit here
}

// Go is fully garbage collected.
// Pointers but no pointer arithmetic.
func learnPointers() (p, q *int) {
    p = new(int)         // new allocates memory.
    // The allocated int is initialized to 0, p is no longer nil.
    s := make([]int, 10) // Allocate 10 ints as a single block of memory.
    s[5] = 24
    r := -44
    return &s[5], &r     // returns two int values.
}

func basicFlowControl() {
    if true {
        fmt.Println("Banana.")
    }
    // Formatting is standardized by the command line command "go fmt."
    if false {
        // Nothing
    } else {
        // Nothing
    }
    // Switch is preffered over chained if statements
    x := 41193.0
    switch x {
    case 0:
    case 1:
    case 41193:
          fmt.Println("Gelato.")
    case 41194:
        // Unreached.
    default:
        // Optional
    }
    // For is the only loop statement in Go
    // Has many alternate forms
    for {        // It is an Infinite loop.
        break    // We are saved by using "break"
        continue // Unreached
    }

    // Use range to iterate over an array, a slice, a string, a map, or a channel.
    // range returns one (channel) or two values (array, slice, string and map).
    for key, value := range map[string]int{"one": 1, "two": 2, "three": 3} {
        // for each pair in the map, print key and value
        fmt.Printf("key=%s, value=%d\n", key, value)
    }

    useDefer()          // Inportant.
    useInterfaces()   // Good stuff coming up!
}

func useDefer() (ok bool) {
    // Deferred statements are executed just before the function returns.
    defer fmt.Println("(SECOND) Deferred statements are executed in reverse.")
    defer fmt.Println("\n(FIRST) Printed first because (LIFO) or ")
    return true
}

// Define infString as interface.
type infString interface {
    String() string
}

// Define stcInt as struct.
type stcInt struct {
    x, y int
}

// Define a method on type stcInt.
// stcInt implements infString.
func (p stcInt) String() string {
    // p is called the "receiver"
    return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func useInterfaces() {
    // p is initialized with values 5 and 6.
    p := stcInt{5, 6}
    fmt.Println(p.String())   // Call String method of p, of type stcInt.
    var i infString           // Declare i of interface type infString.
    i = p                     // Valid because stcInt implements infString
    fmt.Println(i.String())
    fmt.Println(p) // Output same as above.
    fmt.Println(i) // Output same as above.

    useVariadicParams("Kevin", "KingBob", "Dave")
}

// Function with variadic parameters.
func useVariadicParams(myStrings ...interface{}) {
    // Iterate each value of the variadic.
    // The underscore ignores the index argument of the array.
    for _, param := range myStrings {
        fmt.Println("Minion:", param)
    }

    // Pass variadic value as a variadic parameter.
    fmt.Println("All Minions:", fmt.Sprintln(myStrings...))
}
