package main
import "fmt"
import "net/http"
import "os"

func customHandler(w http.ResponseWriter, r *http.Request){
	req_file := "." + r.URL.String() // ex: ./abc.go
	_,err := os.Open(req_file)
	if (err!=nil){
		fmt.Fprintln(w, "File not found")
	} else {
		fs := http.FileServer(http.Dir("."))
		fs.ServeHTTP(w, r)
	}
}

func main(){
	fmt.Println("Launching server...")
	hh := http.HandlerFunc(customHandler)
	http.Handle("/",hh)
	for {
		http.ListenAndServe(":12008",hh)
	}
}