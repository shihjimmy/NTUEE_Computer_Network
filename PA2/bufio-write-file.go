package main 
 
import "fmt" 
import "os" 
import "bufio" 
 
func check(e error) { 
    if e != nil { 
        panic(e) 
    } 
} 
 
func main() { 
	f, err := os.Create("PA2-output.txt") 
	// os.Create() is the API to open a file non-existing yet.
	check(err) 
	defer f.Close() 
/* 	
	Close() is necessary to shut the I/O opened by os.Open() or os.Create(). 
	But one often forgets about closing after hours of coding/debugging. 
	Typing the closing line up as soon the I/O is opened and prepending it with defer prevents the bug entirely.
*/
 
	writer := bufio.NewWriter(f)   // bufio.NewWriter(f) converts a regular I/O (f) to a buffer I/O (writer)
	len, _ := writer.WriteString("This is a test!") 
/*
	1. writer.WriteString() writes a string through the writer to f.
	2. returns two parameters, length of the string and error message. 
	3. _ is used in place of the error message. This is a way to ignore a certain returned parameter 
		if it is not going to be used later anyway.
*/
	
	fmt.Println(len) 
	writer.Flush()  
	// enforce the string temporarily stored on the system memory to the file on the disk.
} 
