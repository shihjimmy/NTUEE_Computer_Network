package main 
 
import "fmt" 
import "bufio" 
import "net" 
import "net/http" 
import "os"

func check(e error) { 
    if e != nil { 
        panic(e) 
    } 
} 

func main(){ 
	fmt.Println("Launching server...") 
    ln, _ := net.Listen("tcp", ":12008") 
    defer ln.Close() 
	
	for {
		conn,_ := ln.Accept()
		reader := bufio.NewReader(conn)
		
		req,err := http.ReadRequest(reader)
		check(err)

		req_file := "." + req.URL.String() // ex: ./abc.go
		file,errr := os.Open(req_file)
		defer file.Close()
		if (errr!=nil || req.URL.String() == "/"){
			fmt.Println("File not found")  
		}else {
			file_info,err2 := os.Stat(req_file)  // take in file path as input
			check(err2)
			fmt.Printf("File size: %d\n", file_info.Size())
		}
		conn.Close()
	}
}