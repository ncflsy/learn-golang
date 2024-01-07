package main

import "fmt"

func main() {
 var a,b string
 fmt.Scanln(&a)
 fmt.Scanln(&b)
 
var temp string
temp = a
a = b
b = temp 
    fmt.Print(a)
    fmt.Print(" ")
    fmt.Print(b)
    
}