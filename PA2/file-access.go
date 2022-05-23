package main

import "fmt"
import "os"

func check(e error){
	if e!=nil{
		panic(e)
	}
}

/*
	(e error) indicates that the function takes in one parameter e of type error.
	If e is not nil, call panic(), which is one of the fundamental APIs defined in the main package. 
	What it does is to show the error message and to force-quit the execution.
*/

func main(){
	f, err := os.Open("hello-world.go")
/*
	os.Open() opens a file, the API returns two parameters 
	and they are assigned to f and err. 
	f is the variable tracking the handle of the file opened. 
	err holds the error message in case of failure.
*/
	check(err)

	word1, word2 := "", "" 
    fmt.Fscanln(f, &word1, &word2) 

/*
	fmt.Fscanln() scans a line from a file.
	Fscanln() reading from a file works just like Scanln() reading from os.Stdin
	
	The first parameter is the file handle. 
	The rest are to hold the strings, separated by a space, in the line. 
	word1 and word2 there will hold only two words in a line.
	
	It will not be general to textual file scanning, where the 
	number of strings is very likely different from line to line.
*/

    fmt.Printf("%s %s\n", word1, word2) 
    for i := 2; i <= 5; i++ { 
        word1, word2 = "", "" 
        fmt.Fscanln(f, &word1, &word2) 
        fmt.Println(word1, word2) 
    } 

/*
	You see Fscanln() is good for files that are well structured, i.e., the number of strings per line is fixed. 
	It unfortunately does not serve all files in general, particularly the textual files.
*/
    f.Close() 
	// close the file
} 
