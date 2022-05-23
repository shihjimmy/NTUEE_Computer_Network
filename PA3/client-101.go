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

	/* 
		The only API being new is net.Dial(). It is used to connect a client to the 
	server. The first input parameter declares the connection type, "tcp" again. 
	The second parameter indicates the IP address and port number the socket 
	connects to. Note the symbol : in the middle, It separates the IP address and 
	the port number.
	*/

    check(errc) 
    defer conn.Close() 
 
    writer := bufio.NewWriter(conn) 
    len, errw := writer.WriteString("Hello World!\n") 
    check(errw) 
    fmt.Printf("Send a string of %d bytes\n", len) 
    writer.Flush() 
 
    scanner := bufio.NewScanner(conn) 
    if scanner.Scan() { 
        fmt.Printf("Server replies: %s\n", scanner.Text()) 
    } 
} 
