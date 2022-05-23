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
    fmt.Println("Launching server...") 
    ln, _ := net.Listen("tcp", ":12008")  
	// 	A socket handle works just like a file handle or the standard I/O

	/* 
		net is the package in which the Listen() function is defined. 
		With the API, the simple server starts a socket and listens for a client request
		the first parameter declares the connection type of the socket
		The second parameter indicates the port number the socket is listening on
		returns two values, the socket handle and the error message
	*/

    conn, _ := ln.Accept() 
    defer ln.Close() 
    defer conn.Close() 

	/* ------------------------
		With ln.Accept(), we call object ln ‘s class method Accept(). The API takes 
		(or waits for) the 1st client request arriving at the listen socket. In turn, it 
		creates another socket conn. conn is dedicated to data transmission between 
		the client and server, so ln can focus on incoming client requests, working 
		pretty much like a call dispatch
	*/

   
    scanner := bufio.NewScanner(conn) 
    message := "" 
    if scanner.Scan() {    //Scan() reads a line from socket conn, without the new line
        message = scanner.Text()  //Text() converts the line (byte stream) to text
        fmt.Println(message) 
    }   
 
    writer := bufio.NewWriter(conn) 
    newline := fmt.Sprintf("%d bytes received\n", len(message)) 
    _, errw := writer.WriteString(newline) 
    check(errw) 
    writer.Flush() 
} 
