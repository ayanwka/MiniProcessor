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

	// проверка артиклей
	for i := 0; i < len(arr); i++ {
		if arr[i] == "a" && ((arr[i+1][0] == 'a') || ((arr[i+1][0] == 'e') || (arr[i+1][0] == 'i') || (arr[i+1][0] == 'o') || (arr[i+1][0] == 'u') || (arr[i+1][0] == 'h'))) {
			arr[i] = "an"
		}
	}

	for i := 0; i < len(arr); i++ {
		switch {
		case strings.HasPrefix(arr[i], ("(up")):
			if arr[i][3] == ')' {
				arr[i-1] = strings.ToUpper(arr[i-1])
				arr[i] = ""
			} else {
				numBefore := strings.TrimSuffix(strings.TrimPrefix(arr[i+1], "(up,"), ")")
				numBefore = strings.TrimSpace(numBefore)
				num, _ := strconv.Atoi(string(numBefore))
				// проверка ошибки в файле input.txt
				if i-num < 0 {
					fmt.Println("index IN FILE IS out of range")
					return
				}
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
				numBefore := strings.TrimSuffix(strings.TrimPrefix(arr[i+1], "(low,"), ")")
				numBefore = strings.TrimSpace(numBefore)
				num, _ := strconv.Atoi(string(numBefore))
				if i-num < 0 {
					fmt.Println("index IN FILE IS out of range")
					return
				}
				for j := i - num; j < i; j++ {
					arr[j] = strings.ToLower(arr[j])
				}
				arr[i] = ""
				arr[i+1] = ""
			}
		case strings.HasPrefix(arr[i], ("(cap")):
			if arr[i][4] == ')' {
				arr[i-1] = strings.Title(arr[i-1])
				arr[i] = ""
			} else {
				numBefore := strings.TrimSuffix(strings.TrimPrefix(arr[i+1], "(cap,"), ")")
				numBefore = strings.TrimSpace(numBefore)
				num, _ := strconv.Atoi(string(numBefore))
				if i-num < 0 {
					fmt.Println("index IN FILE IS out of range")
					return
				}
				for j := i - num; j < i; j++ {
					arr[j] = strings.Title(arr[j])
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
				numBefore := strings.TrimSuffix(strings.TrimPrefix(arr[i+1], "(hex,"), ")")
				numBefore = strings.TrimSpace(numBefore)
				num, _ := strconv.Atoi(string(numBefore))
				if i-num < 0 {
					fmt.Println("index IN FILE IS out of range")
					return
				}
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
				numBefore := strings.TrimSuffix(strings.TrimPrefix(arr[i+1], "(bin,"), ")")
				numBefore = strings.TrimSpace(numBefore)
				num, _ := strconv.Atoi(string(numBefore))
				if i-num < 0 {
					fmt.Println("index IN FILE IS out of range")
					return
				}
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
