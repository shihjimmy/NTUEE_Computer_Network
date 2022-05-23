package main 
 
import "fmt" 
import "net/http"
 
func main() { 
    fmt.Println("Launching server...") 
 
	// HTTP Response by the built-in http.FileServer()
	/* 
		The alternative is to send HTTP response messages exploiting the built-in file server – 
		http.FileServer(). It interprets the incoming HTTP request messages and 
		generates HTTP response messages accordingly
	*/
    http.ListenAndServe(":12008",http.FileServer(http.Dir(".")))
	/*  
		0. http.ListenAndServe 會去監聽 TCP network 的 addr 這個接口,並且呼叫 handler 來處理傳入的 requests,
		這邊的 handler 一般來說會給 nil, 使用預設的 DefaultServeMux.
		1. ListenAndServe() takes in 2 parameters – (1) a port number (e.g., “:8080”) 
		and (2) a function handling the incoming HTTP request message.
		2. What the API does is to (1) start a server listening at the port number and (2) 
		pass the HTTP request message to the handling function, or just “handler”. 
		3. http.FileServer() meant to call the FileServer() API defined in the http 
		package, and FileServer() implements the Golang built-in Web file server. 
		4. FileServer() takes in 1 parameter, the home directory of the Web file 
		server. In the example, it’s specified by http.Dir(“.”)
		5. FileServer(http.Dir(“.”)) meant the built-in server will look from the 
		server’s home directory for the file being requested. 
		6. If the file is not found, FileServer() returns all files in the directory 
		specified by http.Dir(), as a way to suggest alternative files to request.   
		7. If the file is found, FileServer() returns the file to the requested.
	*/
} 
