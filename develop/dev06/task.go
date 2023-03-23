package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Flags struct {
	Fields    string
	Delimiter string
	Separated bool
}

func main() {
	var (
		fFlag = flag.String("f", "", "Выбрать поля")
		dFlag = flag.String("d", "\t", "Использовать другой разделитель")
		sFlag = flag.Bool("s", false, "Только строки с разделителем")
	)
	flag.Parse()
	fl := Flags{Fields: *fFlag, Delimiter: *dFlag, Separated: *sFlag}

	var fileName string
	if flag.Arg(0) != "." {
		fileName = flag.Arg(0)
	} else {
		fileName = "C:\\Desktop\\wb_two\\dev06\\sampleText.txt"
	}
	data := readFile(fileName)
	outData := cutFile(data, fl)
	for _, str := range outData {
		fmt.Println(str)
	}
}

func readFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Ошибка открытия файла!\n", err)
	}
	defer file.Close()

	var text []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return text
}

func cutFile(data []string, flags Flags) []string {
	selectedFields, err := parseFlagField(flags)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	d := "\t"
	if flags.Delimiter != d {
		d = flags.Delimiter
	}

	resStr := []string{}
	for _, line := range data {
		if d != "" && strings.Contains(line, d) {
			words := strings.Split(line, d)

			cutLine := strings.Builder{}

			for _, val := range selectedFields {
				if len(words) >= val {
					cutLine.WriteString(words[val-1])
					cutLine.WriteString(d)
				}
			}
			resStr = append(resStr, strings.TrimSuffix(cutLine.String(), d))

		} else if !flags.Separated {
			resStr = append(resStr, line)
		}
	}
	return resStr
}

func parseFlagField(flags Flags) ([]int, error) {
	tmp := strings.Split(flags.Fields, ",")
	fields := make([]int, len(tmp))
	for i := range tmp {
		num, err := strconv.Atoi(tmp[i])
		if err != nil || num == 0 {
			return nil, fmt.Errorf("Не получилось преобразовать string в int: %v", err)
		}
		fields[i] = num
	}
	return fields, nil
}
