package main

import (
	"fmt"
	"bufio"
	"net"
	"os"
	"io/ioutil"
)
 
func check(e error) { 
    if e != nil { 
        panic(e) 
    } 
} 

func main(){
	// connect to the server
	conn , err := net.Dial("tcp","140.112.42.221:12008")     
	check(err)
	defer conn.Close()

	// prompt the user to input filename
	fmt.Printf("input file name: ")
	filename := ""
	fmt.Scanf("%s",&filename)

	// os get the file's size and send to server 
	filepath := "./" + filename
	f,err2 := os.Stat(filepath)
	check(err2)
	file_size := f.Size()
	fmt.Printf("Send the file size first: %d\n",file_size)
	writer := bufio.NewWriter(conn)
	s := fmt.Sprintf("%d\n",file_size)
	_ , err3 := writer.WriteString(s)
	check(err3)
	writer.Flush()

	// os open the file and send the entire file to the server
	bytes,err5 := ioutil.ReadFile(filepath)
	check(err5)
	str := string(bytes)
	_,err4 := writer.WriteString(str)
	check(err4)
	writer.Flush()

	// get the message back from the server and print out
	scanner_2 := bufio.NewScanner(conn)
	if scanner_2.Scan(){
		fmt.Printf("server replies: %s\n",scanner_2.Text())
	}

}
