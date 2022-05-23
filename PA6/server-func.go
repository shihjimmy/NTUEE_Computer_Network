package main 
 
import "fmt" 
import "bufio" 
import "net" 
 
 
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