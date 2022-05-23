package main 
 
import "fmt" 
import "bufio" 
import "net" 
import "net/http" 
 
func check(e error) { 
    if e != nil { 
        panic(e) 
    } 
} 
 
func main() { 
    fmt.Println("Launching server...") 
    ln, _ := net.Listen("tcp", ":12008") 
    defer ln.Close() 
    conn, _ := ln.Accept() 
    defer conn.Close() 
 
    reader := bufio.NewReader(conn) 
    req, err := http.ReadRequest(reader) 
    check(err) 
    fmt.Printf("Method: %s\n", req.Method) 
 
	/*
		func Fprintf(w io.Writer, format string, a ...any)(n int, err error)
		Fprintf 根據格式說明符格式化並寫入 w。它返回寫入的字節數和遇到的任何寫入錯誤。
	*/

	// HTTP Response as a string
    fmt.Fprintf(conn, "HTTP/1.1 404 Not Found\r\n") 
    fmt.Fprintf(conn, "Date: ...\r\n") 
    fmt.Fprintf(conn, "\r\n") 
    fmt.Fprintf(conn, "File not found\r\n") 
    fmt.Fprint(conn, "\r\n") 
}