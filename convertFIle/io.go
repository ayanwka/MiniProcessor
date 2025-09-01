package Convert

import "os"

func ReadFile(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return data
}

func WriteFile(filenameRead, filenameWrite string) {
	FileTransform(ReadFile(filenameRead), filenameWrite)
}
