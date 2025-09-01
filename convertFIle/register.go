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
		switch {
		case strings.HasPrefix(arr[i], ("(up")):
			if arr[i][3] == ')' {
				arr[i-1] = strings.ToUpper(arr[i-1])
				arr[i] = ""
			} else {
				numBefore := arr[i+1][0]
				num, _ := strconv.Atoi(string(numBefore))
				for j := i - num; j < i; j++ {
					arr[j] = strings.ToUpper(arr[j])
				}
				arr[i] = ""
				arr[i+1] = ""
			}
		case strings.HasPrefix(arr[i], ("(low")):
			if arr[i][4] == ')' {
				arr[i-1] = strings.ToLower(arr[i-1])
				arr[i] = ""
			} else {
				numBefore := arr[i+1][0]
				num, _ := strconv.Atoi(string(numBefore))
				for j := i - num; j < i; j++ {
					arr[j] = strings.ToLower(arr[j])
				}
				arr[i] = ""
				arr[i+1] = ""
			}
		case strings.HasPrefix(arr[i], ("(cap")):
			if arr[i][4] == ')' {
				arr[i-1] = strings.ToTitle(arr[i-1])
				arr[i] = ""
			} else {
				numBefore := arr[i+1][0]
				num, _ := strconv.Atoi(string(numBefore))
				for j := i - num; j < i; j++ {
					arr[j] = strings.ToTitle(arr[j])
				}
				arr[i] = ""
				arr[i+1] = ""
			}
		case strings.HasPrefix(arr[i], ("(hex")):
			if arr[i][4] == ')' {
				//ставим _, т.к convertint возвращает 2 значения. string, error
				arr[i-1], _ = ConvertInt(arr[i-1], 16, 10)
				arr[i] = ""
			} else {
				numBefore := arr[i+1][0]
				num, _ := strconv.Atoi(string(numBefore))
				for j := i - num; j < i; j++ {
					arr[j], _ = ConvertInt(arr[j], 16, 10)
				}
				arr[i] = ""
				arr[i+1] = ""
			}
		case strings.HasPrefix(arr[i], ("(bin")):
			if arr[i][4] == ')' {
				ConvertInt(arr[i-1], 2, 10)
				arr[i-1], _ = ConvertInt(arr[i-1], 2, 10)
				arr[i] = ""
			} else {
				numBefore := arr[i+1][0]
				num, _ := strconv.Atoi(string(numBefore))
				for j := i - num; j < i; j++ {
					arr[j], _ = ConvertInt(arr[j], 2, 10)
				}
				arr[i] = ""
				arr[i+1] = ""
			}
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
