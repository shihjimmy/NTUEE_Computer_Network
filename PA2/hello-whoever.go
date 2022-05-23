package main 
 
import "fmt" 
import "os"
 
func main() { 
     fmt.Printf("Who's there?\n") 
     text := ""   
	 // declares the variable text and assigns an empty string to it
	 //In Go,to assign value to an existing variable, just say =. := is to declare and assign at the same time
	 //the compiler identifies the data type automatically,by looking at the initial value
     
	 fmt.Scanf("%s", &text) 
	 // %s means string
 
     fmt.Printf("Hello, %s\n", text) 
     fmt.Println("Hello,", text) 
	 // Println() is pronounced print line, as it enforces a new line after execution
     fmt.Fprintf(os.Stdout, "Hello, %s\n", text) 

	 /* 
	 Fprintf() means printing to a file in fact
	 In Unix, a file is also an I/O.
	 Therefore, writing to the display is equivalent of writing to a file at os.Stdout.
	 
	 fmt.Fprintf(os.Stdout,) is equivalent of fmt.Printf(). 
	 fmt.Fscanf(os.Stdin,) is equivalent of fmt.Scanf().
	 
	 */
} 
