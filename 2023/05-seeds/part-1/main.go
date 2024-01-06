package main

import "fmt"
import "os"

func main() {
	args := os.Args
	fPath := args[1]
	data, _ := os.ReadFile(fPath)
	content := string(data)
	fmt.Println(content)
}
