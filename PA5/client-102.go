package main 
 
import "fmt" 
import "bufio" 
import "net" 
 
 
func check(e error) { 
    if e != nil { 
        panic(e) 
    } 
} 
 
func main() { 
    conn, errc := net.Dial("tcp", "127.0.0.1:12008") 
    check(errc) 
    defer conn.Close() 
 
    writer := bufio.NewWriter(conn) 
    len, errw := writer.WriteString("Hello World!\n") 
    check(errw) 
    fmt.Printf("Send a string of %d bytes\n", len) 
    writer.Flush() 
 
    reader := bufio.NewReader(conn) 

	/*
		bufio.NewReader(conn) works just like bufio.NewScanner(). 
		In that, NewReader() creates a Reader object that wraps around conn to allow the 
		socket connection access to APIs available to the Reader type
	*/

    message, errr := reader.ReadString('\n') 

	/* 
		reader.ReadString(\n) reads a string delimited by the character in the (). 
		In effect, this call reads a line, including the \n.
		Giving it, i.e., the delimiter, an empty space, the call will read a word. 
		Text extraction/processing is even easier with Reader than with Scanner
	*/

    check(errr) 
    fmt.Printf("Server replies: %s", message) 
} 
