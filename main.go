package main

import (
	"fmt"
	"os"
	Convert "test_game1/convertFIle"
)

func main() {
	// os.Args принимает аргументы прямо из консоли
	// go run . input.txt output.txt
	file1 := os.Args[1]
	file2 := os.Args[2]
	// 3, потому что os.Args всегда включает имя программы как первый элемент
	if len(os.Args) != 3 {
		fmt.Println("enter: go run . file_1 file_2")
		return
	}
	Convert.WriteFile(string(file1), string(file2))
}
