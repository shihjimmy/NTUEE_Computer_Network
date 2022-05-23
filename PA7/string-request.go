package main 
 
import "fmt" 
import "bufio" 
import "net" 
import "strings" 
 
func check(e error) { 
    if e != nil { 
        panic(e) 
    } 
} 

//	HTTP Request as a string
func main() { 
    fmt.Println("Launching server...") 
    ln, _ := net.Listen("tcp", ":12008") 
    defer ln.Close() 
    conn, _ := ln.Accept() 
    defer conn.Close() 
    reader := bufio.NewReader(conn) 
    for { 
        req, err := reader.ReadString('\n') 
        check(err) 
        if req == "\r\n" { 
            break 
        } 
        tokens := strings.Split(req, " ")   
	//	strings.Split() returns an array of substrings (tokens) as the result of the split.
        for i := range tokens { 
            fmt.Println(tokens[i]) 
        } 
    } 
} 
