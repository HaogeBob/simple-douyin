package util

import "fmt"

func GetFileUrl(fileName string) string {
	base := fmt.Sprintf("http://127.0.0.1:8080/static/%s", fileName)
	return base
}
