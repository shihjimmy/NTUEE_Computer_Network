package main 
 
import "fmt" 
import "bufio" 
import "net" 
import "net/http" 
/*
	net/http is the package where the http-related data struct and APIs are 
	defined. Contained in the package is also a Web file server itself.
*/ 
 
func check(e error) { 
    if e != nil { 
        panic(e) 
    } 
} 
 
//	HTTP Request as a Request 
func main() { 
    fmt.Println("Launching server...") 
    ln, _ := net.Listen("tcp", ":12008") 
    defer ln.Close() 
    conn, _ := ln.Accept() 
    defer conn.Close() 
 
    reader := bufio.NewReader(conn) 
    req, err := http.ReadRequest(reader) 
    check(err) 

	/* 
		http.ReadRequest() is a special bufio reader API defined in net/http 
	(not in bufio). It reads from the socket, process the textual input, and store 
	the information as a Request object, whose struct is defined in 
	net/http/request.go.
		the first value req is the Request object containing the interpreted HTTP request. 
	the second value err carries the error code.
	*/

    fmt.Printf("Method: %s\n", req.Method) 
    fmt.Printf("Host: %s\n", req.Host) 
    fmt.Printf("User-Agent: %s\n", req.UserAgent()) 
} 
