package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	wordsArr := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	res := searchAnagram(&wordsArr)
	for k, val := range *res {
		fmt.Printf("%q: %v\n", k, val)
	}
}

func searchAnagram(words *[]string) *map[string][]string {
	anagrams := make(map[string][]string)

	usesWords := map[rune][]string{}
	usesKeys := map[int32]string{}

	toLowerSlice(words)

	for i := 0; i < len(*words); i++ {
		countRuneInWord := int32(0)
		for _, val := range (*words)[i] {
			countRuneInWord += val
		}
		if _, ok := usesKeys[countRuneInWord]; !ok {
			usesKeys[countRuneInWord] = (*words)[i]
			usesWords[countRuneInWord] = append(usesWords[countRuneInWord], (*words)[i])
		} else {
			usesWords[countRuneInWord] = append(usesWords[countRuneInWord], (*words)[i])
		}
	}

	for _, val := range usesWords {
		sort.Strings(val)
	}

	for k, val := range usesKeys {
		if len(usesWords[k]) > 1 {
			anagrams[val] = usesWords[k]
		}
	}
	return &anagrams
}

func toLowerSlice(words *[]string) {
	for i := 0; i < len(*words); i++ {
		(*words)[i] = strings.ToLower((*words)[i])
	}
}
