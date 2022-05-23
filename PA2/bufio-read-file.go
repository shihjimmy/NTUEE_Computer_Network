package main 
 
import "fmt" 
import "os" 
import "bufio" 

/*
	You see Fscanln() is good for files that are well structured, i.e., the number of strings per line is fixed. 
	It unfortunately does not serve all files in general, particularly the textual files. 
	This leads us to bufio, a package to treat I/Os as general byte streams.
*/
 
func check(e error) { 
    if e != nil { 
        panic(e) 
    } 
} 
 
func main() { 

// 	os.Open() require a absolute path
	f, err := os.Open("C:\\Users\\user\\Desktop\\computer networks\\programming_assignments\\PA2\\bufio-read-file.go") 
	check(err) 
 
	scanner := bufio.NewScanner(f) 
/* 
	NewScanner() is an API defined in the bufio package. 
	It converts a regular I/O (f in this case) to a buffer I/O – a Scanner object in this case.
*/
	for scanner.Scan() { 
 		fmt.Println(scanner.Text()) 
	} 
/*
	scanner.Scan() scans from f. The default is to scan one line at a time.
	One can configure it to scan one word or one byte at a time.(check out the go document)
	When the scanner.Scan() reaches the end of file, it returns false (true otherwise).
	scanner.Text() converts a byte stream to a string so it can be printed using fmt.Println().
*/

 	f.Close() 
} 
