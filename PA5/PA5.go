package main

import "fmt"
import "net"
import "bufio"
import "os"
import "strings"
import "strconv"

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
		defer conn.Close()
		
		new_file,err := os.Create("whatever.txt")
		check(err)
		defer new_file.Close()
		filewriter := bufio.NewWriter(new_file)
	
		reader := bufio.NewReader(conn)
		first_message,errr := reader.ReadString('\n')
		check(errr)
		first_message = strings.TrimRight(first_message,"\n")
		file_size, _ := strconv.Atoi(first_message)

		line_count := 1
		total_size := 0
		count_size := 0
		for {
			message, err2 := reader.ReadString('\n')
			count_size += len(message)
			check(err2)
			message = fmt.Sprintf("%d ",line_count) + message
			line_count += 1
			total_size += len(message)

			_,err3 := filewriter.WriteString(message)
			check(err3)
		
			if count_size==file_size {
				break
			}
		}
		filewriter.Flush()

		response := fmt.Sprintf("Upload file size: %d\nOutput file size: %d\n",file_size,total_size)
		fmt.Printf(response)
	
		reply := fmt.Sprintf("%d bytes received, %d bytes file generated\n",file_size,total_size)	
		writer := bufio.NewWriter(conn)
		_,err4 := writer.WriteString((reply))
		check(err4)
		writer.Flush()
	}
}