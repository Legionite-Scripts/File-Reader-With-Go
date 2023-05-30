package main

import (
	"fmt"
	"os"
	"errors"
)

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		// Error Handling if file doesn't exist
		fmt.Print(errors.New("File does not exist"))
		return
	}
	defer file.Close() //deferring the flile.Close() method till the function ends
	// to get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	// to 	read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}
