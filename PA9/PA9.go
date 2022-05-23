package main

import "fmt"
import "net/http"
import "os"

func CustomHandler(w http.ResponseWriter, r *http.Request){
 // check if the requested file is in this folder.
	req_file := "." + r.URL.String() // ex: ./abc.go
	_,err := os.Open(req_file)
	if (err!=nil || os.IsNotExist(err)){
		error404Handler(w,r)
		return
	} else {
		http.ServeFile(w,r,req_file)
	}
}

func error404Handler(w http.ResponseWriter, r *http.Request) { 
    http.Error(w, "File not found", http.StatusNotFound) 
}

func main(){
	fmt.Println("Launching server...")
 	hh := http.HandlerFunc(CustomHandler)
 	http.Handle("/",hh)
 	http.ListenAndServeTLS(":12008", "server.cer","server.key", nil)
}