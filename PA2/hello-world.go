package main
// With package main, the compiler includes the most basic, minimum set of APIs

import "fmt"
// import indicates the additional API sets to include
// fmt package contains APIs  
// that generate output or take input of multiple formats for a variety of system I/Os

func main() {
	fmt.Printf("hello world!\n")
}