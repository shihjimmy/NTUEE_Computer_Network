package main 
 
import "fmt" 
import "net/http"  
 
func helloHandler(w http.ResponseWriter, r *http.Request) { 
    fmt.Fprintln(w, "Hello, world!") 
} 
 
func main() { 
    fmt.Println("Launching server...") 
 
    hh := http.HandlerFunc(helloHandler) 
    http.Handle("/hello", hh) 
    fs := http.FileServer(http.Dir(".")) 
    http.Handle("/", http.StripPrefix("/", fs)) 
    http.ListenAndServeTLS(":12008", "server.cer", "server.key", nil) 
	/* 
		The only API new is http.ListenAndServeTLS(). It is the equivalent of 
	http.ListenAndServe() to start secure communication. The usage is slightly 
	different as it takes the public and private key file as inputs as well.
	*/
} 
