package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	// cрез, где: os.Args[0] — имя программы, os.Args[1] — первый аргумент (путь к файлу).
	if len(os.Args) != 2 {
		fmt.Println("Использование: go run main.go <путь_к_файлу>")
		os.Exit(1)
	}
	filePath := os.Args[1]

	// обязательно проверяем err != nil — иначе программа упадёт при отсутствии файла
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Ошибка чтения файла %q: %v\n", filePath, err)
		os.Exit(1)
	}

	// []byte в string.
	text := string(data)

	// русские буквы UTF-8 = 2 байта, а не 1   len("Привет") вернёт 12
	// utf8.RuneCountInString("Привет") — 6
	// количество символов (рун, а не байт).
	charCount := utf8.RuneCountInString(text)

	// В файле из N строк обычно N-1 символов '\n' — добавляем 1 к количеству '\n'.
	lineCount := strings.Count(text, "\n") + 1

	// разбивает строку на срез подстрок по любому количеству символов (пробел, \t, \n, \r)
	words := strings.Fields(text)
	wordCount := len(words)

	fmt.Printf("__ Анализ файла: %s __\n", filePath)
	fmt.Printf("Количество символов: %d\n", charCount)
	fmt.Printf("Количество строк: %d\n", lineCount)
	fmt.Printf("Количество слов: %d\n", wordCount) 
}