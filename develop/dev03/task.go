package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===
Отсортировать строки (man sort)
Основное

Поддержать ключи
-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное
Поддержать ключи
-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Flags struct {
	K int
	N *bool
	R *bool
	U *bool
}

func main() {
	var (
		collumn int
		num     *bool
		revSort *bool
		uniqStr *bool
	)
	flag.IntVar(&collumn, "k", -1, "Количество столбцов для сортировки")
	num = flag.Bool("n", false, "Сортировка по числовому столбцу, обязательно указывать ключ -k")
	revSort = flag.Bool("r", false, "Сортировка в обратном порядке, обязательно указывать ключ -k")
	uniqStr = flag.Bool("u", false, "Удалить повторения")

	flag.Parse()

	fl := Flags{K: collumn, N: num, R: revSort, U: uniqStr}

	var fileName string
	if flag.Arg(0) != "" {
		fileName = flag.Arg(0)
	} else {
		fileName = "C:\\Desktop\\wb_two\\develop\\dev03\\sampleText.txt"
	}

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

	sortFile(&text, fl)

	if err := os.WriteFile("C:\\Desktop\\wb_two\\develop\\dev03\\resultText.txt",
		[]byte(strings.Join(text, "\n")), os.ModePerm); err != nil {
		fmt.Println("Ошибка записи в файл\n", err)
	}
}

func sortFile(data *[]string, flags Flags) {
	if *flags.U {
		remDup(data)
	}

	if flags.K != -1 {
		flags.K--
		if *flags.N {
			sortColumnNum(data, flags.K)
		} else {
			sortColumnString(data, flags.K)
		}
	}

	if *flags.R {
		reverseStringSlice(data)
	}
}

func remDup(data *[]string) {
	setStr := make(map[int]struct{})
	setInd := make(map[int]struct{})
	for ind, str := range *data {
		hash := hashSumBytes(str)
		if _, ok := setStr[hash]; ok {
			setInd[ind] = struct{}{}
		} else {
			setStr[hash] = struct{}{}
		}
	}
	delCount := 0
	for k := range setInd {
		k -= delCount
		if k == len(*data)-1 {
			*data = (*data)[:k-1]
		} else {
			*data = append((*data)[:k], (*data)[k+1:]...)
		}
		delCount++
	}
}

func hashSumBytes(str string) int {
	sum := 0
	for ind, ch := range str {
		sum += ind * int(ch)
	}
	return sum
}

func sortColumnString(data *[]string, column int) {
	mCol := make(map[string]int)
	for ind, str := range *data {
		words := strings.Split(str, " ")
		if column >= len(words) {
			fmt.Println("Неверно указан столбец!")
			os.Exit(1)
		}
		words[column] += strconv.Itoa(ind)
		mCol[words[column]] = ind
	}

	keys := make([]string, 0, len(mCol))
	for k := range mCol {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	retData := []string{}
	for _, val := range keys {
		retData = append(retData, (*data)[mCol[val]])
	}

	*data = retData
}

func sortColumnNum(data *[]string, column int) {
	type structForSort struct {
		col float64
		ind int
	}
	answerStruct := []structForSort{}
	for ind, str := range *data {
		words := strings.Split(str, " ")
		if column >= len(words) {
			fmt.Println("Неверно указан столбец!")
			os.Exit(1)
		}
		word := strings.Split(words[column], ".")
		var num float64
		if len(word) > 1 {
			iPart, _ := strconv.Atoi(word[0])
			zPart, _ := strconv.Atoi(word[1])
			num = float64(iPart) + float64(zPart/len(word[1]))
		} else {
			numInt, _ := strconv.Atoi(word[0])
			num = float64(numInt)
		}
		answerStruct = append(answerStruct, structForSort{num, ind})
	}

	sort.Slice(answerStruct, func(i, j int) (less bool) {
		return answerStruct[i].col < answerStruct[j].col
	})

	retData := []string{}
	for _, val := range answerStruct {
		retData = append(retData, (*data)[val.ind])
	}

	*data = retData
}

func reverseStringSlice(data *[]string) {
	i := 0
	j := len(*data) - 1
	for i < j {
		(*data)[i], (*data)[j] = (*data)[j], (*data)[i]
		i++
		j--
	}
}
