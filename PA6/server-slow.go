package main 
 
import "fmt" 
import "bufio" 
import "net" 
import "time" 
 
 
func check(e error) { 
    if e != nil { 
        panic(e) 
    } 
} 
 
func handleConnection (c net.Conn) { 
    reader := bufio.NewReader(c) 
    message, errr := reader.ReadString('\n') 
    check(errr) 
    fmt.Printf("%s", message) 
 
    writer := bufio.NewWriter(c) 
    newline := fmt.Sprintf("%d bytes received\n", len(message)) 
    _, errw := writer.WriteString(newline) 
    check(errw) 
    writer.Flush() 
 
    time.Sleep(5 * time.Second) 
	/* 
		time.Sleep() is an API that enforces a sleep time, i.e., adding a pause of a 
	specific duration to the process.
		time.Second is a constant that counts the time in seconds.   
		What’s happening here is that the 1st go run client-102.go takes the 
	server into the 5 second wait. When the 2nd go run client-102.go is fired, 
	the client’s net.Dial() needs to wait until the 5-second wait from the 1st fire 
	expires and the server returns to the for loop to ln.Accept() the client’s 
	net.Dial(). 
	*/
} 

func main() { 
    fmt.Println("Launching server...") 
    ln, _ := net.Listen("tcp", ":12008") 
    defer ln.Close() 
 
    i := 1 
    for {   //infinite loop
        conn, _ := ln.Accept() 
        defer conn.Close() 
 
        fmt.Printf("%d ", i) 
        handleConnection(conn) 
        i++ 
    } 
}