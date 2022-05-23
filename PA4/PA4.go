package main

import "fmt"
import "io"
import "net"
import "bufio"
import "os"
import "strconv"   // turn string to int or turn int to string
import "strings"

func check(e error) { 
    if e != nil { 
        panic(e) 
    } 
} 

func main(){
	fmt.Println("Server is launched...")
	ln, _ := net.Listen("tcp",":12008")
	conn,_:= ln.Accept()
	defer conn.Close()
	defer ln.Close()

	reader := bufio.NewReader(conn)
	f,err2 := os.Create("whatever.txt")
	check(err2)
	defer f.Close()
	writer := bufio.NewWriter(f)

	firstmsg, _ := reader.ReadString('\n')  //get file size in string format
	firstmsg = strings.TrimRight(firstmsg,"\n") //delete the right side \n to get right size
	original_size, _ := strconv.Atoi(firstmsg) //get file size in int format
	line_count := 1
	file_size := 0
	add_size := 0

	for {  //infinite loop
		message, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}
		check(err)
		file_size += len(message)
		add_size += ( len(strconv.Itoa(line_count)) +1)
		
		message = fmt.Sprintf("%d ",line_count) + message
		_,err4 := writer.WriteString(message)
		check(err4)
		line_count += 1
		
		if file_size==original_size {
			break
		}
	
	}
	writer.Flush()
	
	response := fmt.Sprintf("original file size: %d\tnew file size: %d\n",original_size, file_size+add_size)
	reply := bufio.NewWriter(conn)
	_, err5 := reply.WriteString(response)
	check(err5)
	reply.Flush()
	
}