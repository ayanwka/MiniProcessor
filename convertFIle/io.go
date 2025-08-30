package Convert

import "os"

func ReadFile() []byte {
	filename := "text.txt"
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return data
}

func WriteFile() {
	filename := "output.txt"

	FileTransform(ReadFile(), filename)
}
