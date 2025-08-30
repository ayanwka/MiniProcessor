package Convert

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FileTransform(file []byte, name string) {
	ByteToString := string(file)
	arr := strings.Fields(ByteToString)
	for i := 0; i < len(arr); i++ {
		switch arr[i] {
		case "(up)":
			arr[i-1] = strings.ToUpper(arr[i-1])
			arr[i] = ""
		case ".up":
			num, _ := strconv.Atoi(arr[i+1])
			for j := i - num; j < i; j++ {
				arr[j] = strings.ToUpper(arr[j])
			}
			arr[i] = ""
			arr[i+1] = ""
		case "(low)":
			arr[i-1] = strings.ToLower(arr[i-1])
			arr[i] = ""
		case ".low":
			num, _ := strconv.Atoi(arr[i+1])
			for j := i - num; j < i; j++ {
				arr[j] = strings.ToLower(arr[j])
			}
			arr[i] = ""
			arr[i+1] = ""
		case "(cap)":
			arr[i-1] = strings.Title(arr[i-1])
			arr[i] = ""
		case ".cap":
			num, _ := strconv.Atoi(arr[i+1])
			for j := i - num; j < i; j++ {
				arr[j] = strings.Title(arr[j])
			}
			arr[i] = ""
			arr[i+1] = ""
		case "(hex)":
			//ставим _, т.к convertint возвращает 2 значения. string, error
			arr[i-1], _ = ConvertInt(arr[i-1], 16, 10)
			arr[i] = ""
		case ".hex":
			num, _ := strconv.Atoi(arr[i+1])
			for j := i - num; j < i; j++ {
				arr[j], _ = ConvertInt(arr[j], 16, 10)
			}
			arr[i] = ""
			arr[i+1] = ""
		case "(bin)":
			ConvertInt(arr[i-1], 2, 10)
			arr[i-1], _ = ConvertInt(arr[i-1], 2, 10)
			arr[i] = ""
		case ".bin":
			num, _ := strconv.Atoi(arr[i+1])
			for j := i - num; j < i; j++ {
				arr[j], _ = ConvertInt(arr[j], 2, 10)
			}
			arr[i] = ""
			arr[i+1] = ""
		}
	}

	ArrToString := strings.Join(arr, " ")
	StringToByte := []byte(ArrToString)

	err := os.WriteFile(name, []byte(StringToByte), 0600)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(StringToByte))
}
