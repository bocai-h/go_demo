package main

import (
 "fmt"
 "encoding/csv"
 "io"
 "os"
 )

func main() {
	// file, err := os.Open("/home/bocai/Downloads/part-00000-5205cd07-3aac-4b8c-9403-e730ff406b19-c000.csv")
    file, err := os.Open("./test2.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.LazyQuotes = true

	for {
        record, err := reader.Read()
        if err == io.EOF {
        	break
        }else if err != nil {
        	fmt.Println("Error:", err)
        	return
        }
        fmt.Println(record)
	}
}