package main

import (
	"os"
	Convert "test_game1/convertFIle"
)

func main() {
	file1 := os.Args[1]
	file2 := os.Args[2]
	Convert.WriteFile(string(file1), string(file2))
}
