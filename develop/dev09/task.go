package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	url := flag.Arg(0)
	page := wget(url)
	writeFile(url, page)
}

func writeFile(url string, strs []byte) {
	var path string
	url = strings.TrimPrefix(url, "https://")
	domens := strings.Split(url, "/")
	file := domens[len(domens)-1]
	path = "C:\\Desktop\\wb_two\\develop\\dev09\\" + file + ".html"

	if err := os.WriteFile(path,
		strs, os.ModePerm); err != nil {
		fmt.Println("Ошибка записи в файл\n", err)
	}
}

func wget(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return body
}
