package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	home := os.Getenv("HOME")
	fmt.Println(home)

	err := os.Mkdir("example_dir", 0755)
	if err != nil {
		fmt.Println("Error Creating Directory", err)
		return
	}

	// create a file
	file, err := os.Create("example_dir/example_file.txt")
	if err != nil {
		fmt.Println("Error creating file", err)
		return
	}

	_, err = file.WriteString("Hello this is OS Package\n")
	if err != nil {
		fmt.Println("Error writing to file", err)
		return
	}
	err = file.Close()
	if err != nil {
		fmt.Println("Error closing file ", err)
		return
	}

	file, err = os.Open("example_dir/example_file.txt")
	if err != nil {
		fmt.Println("Error opening file ", err)
		return
	}

	defer file.Close()
	content := make([]byte, 1024)

	n, err := file.Read(content)

	if err != nil && err != io.EOF {
		fmt.Println("Error reading file", err)
		return
	}
	fmt.Println("File content :- ", string(content[:n]))

	// Remove a file
	err = os.Remove("example_dir/example_file.txt")
	//
	// if err != nil {
	// 	fmt.Println("Error removing file ", err)
	// 	return
	// }
	//
	// Remove directory
	// err = os.Remove("example_dir")

	// if err != nil {
	// 	fmt.Println("Error removing directory ", err)
	// 	return
	// }

	fmt.Println("Done")

}
