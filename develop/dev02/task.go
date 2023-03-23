package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	input := "qwerty3test\\6\\sec\\9qwe\\\\4"
	output, err := unpack(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(input)
	fmt.Println(output)
}

func unpack(str string) (string, error) {
	answer := strings.Builder{}
	for i, sym := range str {
		if unicode.IsDigit(sym) {
			if i == 0 {
				return "", errors.New("некорректная строка")
			}
			ch := str[i-1]
			if unicode.IsDigit(rune(ch)) {
				continue
			}
			if ch == '\\' {
				if str[i-2] != '\\' {
					num := searchNum(str, i+1)
					if num == 0 {
						answer.WriteString(string(str[i]))
					} else {
						for i := 0; i < num; i++ {
							answer.WriteString(string(str[i]))
						}
					}
					continue
				}
			}
			num := searchNum(str, i)
			for i := 0; i < num; i++ {
				answer.WriteString(string(ch))
			}
		} else if str[i] == '\\' {
			continue
		} else if i == len(str)-1 || !unicode.IsDigit(rune(str[i+1])) {
			answer.WriteString(string(str[i]))
		}
	}

	return answer.String(), nil
}

func searchNum(str string, startInd int) int {
	count := 0
	for _, n := range str[startInd:] {
		if unicode.IsDigit(n) {
			count++
		} else {
			break
		}
	}
	num, _ := strconv.Atoi(str[startInd : startInd+count])
	return num
}
