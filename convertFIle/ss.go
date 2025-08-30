package Convert

import (
	"strconv"
)

func ConvertInt(val string, base, toBase int) (string, error) {
	// ParseInt|ParseFloat для конвертации строки в число. 64 это вид (int32, int 64..)
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	// возвращает конвертированное значение (в string) и проверку на ошибку. где i это изначальное слово, tobase это система в которую переводим
	return strconv.FormatInt(i, toBase), nil
}
