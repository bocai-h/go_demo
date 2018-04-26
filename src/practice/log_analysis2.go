package main 

import (
  "bufio"
  "encoding/json"
  "fmt"
  "os"
  "regexp"
  "strconv"
  "strings"
)

var digitsRegexp = regexp.MustCompile(`(IOS|ADR)-\d+-\d+-\d+`)

func main() {
	fileHandle, _ := os.Open("./test.csv")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		fmt.Println("====xxxx====")
		fmt.Println(digitsRegexp.FindString(fileScanner.Text()))
		metaData := strings.Split(fileScanner.Text(), "\t")
		s, err := strconv.Unquote(metaData[1])
		if err!=nil {
			fmt.Printf("err:%+v\n", err)
			return
		}
		var logs []string
		if err := json.Unmarshal([]byte(s), &logs); err != nil {
			fmt.Println("err:%+v\n")
		}
		fmt.Println("=====####=====")
		for _, l := range logs {
			fmt.Println("------------------")
			fmt.Println(l)
		}
	}
}