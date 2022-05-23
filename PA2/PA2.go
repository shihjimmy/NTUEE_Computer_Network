package main

import (
	"fmt"
	"os"
	"bufio"
)

func check(err error){
	if err!=nil {
		panic(err)
	}
}

func main(){
	// Type your Input file name
	fmt.Printf("Input file name: ")
	input_name := ""
	fmt.Scanf("%s", &input_name)
	f_input, err := os.Open(input_name)
	check(err)
	defer f_input.Close()

	// Type your Output file name
	fmt.Printf("Output file name: ")
	output_name := ""
	fmt.Scanf("%s",&output_name)
	f_output, err := os.Create(output_name)
	check(err)
	defer f_output.Close()

	//Write the output file
	scanner := bufio.NewScanner(f_input)
	writer := bufio.NewWriter(f_output)

	count := 1
	text := ""
	for scanner.Scan() {
		text = scanner.Text()
		new_text := fmt.Sprintf("%d ",count) + text + "\n"

		writer.WriteString(new_text)
		count = count + 1
		writer.Flush()
	}

}