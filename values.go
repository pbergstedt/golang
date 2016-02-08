// Go has various value types including strings,
// integers, floats, booleans, etc. Here are a few
// basic examples.

package main

import "fmt"
import "time"

func main() {

    // Strings, which can be added together with `+`.
    fmt.Println("go" + "lang")

    // Integers and floats.
    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    // Booleans, with boolean operators as you'd expect.
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)

    // `var` declares 1 or more variables.
    // var a string = "initial"
    // fmt.Println(a)

    var pjb string = "PJB"
    fmt.Println(pjb)
    fmt.Println("My initials are:", pjb)

    // You can declare multiple variables at once.
    var b, c int = 1, 2
    fmt.Println(b, c)

    // Go will infer the type of initialized variables.
    var d = true
    fmt.Println(d)

    // Variables declared without a corresponding
    // initialization are _zero-valued_. For example, the
    // zero value for an `int` is `0`.
    var e int
    fmt.Println(e)

    // The `:=` syntax is shorthand for declaring and
    // initializing a variable, e.g. for
    // `var f string = "short"` in this case.
    f := "short"
    fmt.Println(f)

    var li int = 1
    for li <= 7 {
        fmt.Println(li)
        li = li + 1
    }

    if 7%2 == 0 {
            fmt.Println("7 is even")
        } else {
            fmt.Println("7 is odd")
    }

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("it's before noon")
    default:
        fmt.Println("it's after noon")
    }
    fmt.Println(t)

    var a [5]int
    fmt.Println("initial array", a)

    a[2] = 22
    fmt.Println("set possition 3:", a)
    fmt.Println("get the 3rd digit:", a[2])

    ab := [5]int{45005, 45402, 45236, 88901, 20015}
    fmt.Println("dcl:", ab)
    fmt.Println("lenght of array:", len(ab))

    n := map[string]int{"Dayton": 45402, "Middletown": 45042, "Franklin": 45005}
        fmt.Println("map:", n)

    v1 := n["Dayton"]
    fmt.Println("v1: ", v1)

}
