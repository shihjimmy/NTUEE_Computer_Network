package main 
 
import "fmt" 
import "net/http"  
 
// HTTP Response by the custom http.FileServer()
/*
	there are ways to customize the Web server and the file server as well. 
	The Web server is designed so that a client can send HTTP requests with specific URL prefixes for 
	custom functionalities.
*/
func helloHandler(w http.ResponseWriter, r *http.Request) { 
    fmt.Fprintln(w, "Hello, world!") 
} 
 
func main() { 
    fmt.Println("Launching server...") 
 
    hh := http.HandlerFunc(helloHandler)    
	// hh is obtained by adapting a programmer-defined function to a handler function 
    http.Handle("/hello", hh)
	// 1. http.Handle() is the key API that associates a prefix to its handler and 
	// 	  inserts the prefix-handler entry to the DefaultServeMux. 
	// 2. The handler-Response.go server handles the /hello command and works as well as a Web file server.
	// 3. As a result, if the URL starts with /hello , helloHandler will be called 
    fs := http.FileServer(http.Dir(".")) 
    http.Handle("/", http.StripPrefix("/", fs))
	// if the URL starts with / , StripPrefix("/", fs) will be called.
    http.ListenAndServe(":12008", nil) 
} 
