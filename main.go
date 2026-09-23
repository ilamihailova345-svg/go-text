package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type WordStat struct {
	Word  string
	Count int
}

type TextStats struct {
	CharCount int
	WordCount int
	LineCount int
	wordFreq  map[string]int
}

func analyzeFile(path string) (TextStats, error) {
	var stats TextStats

	file, err := os.Open(path)
	if err != nil {
		return stats, err
	}
	defer file.Close()

	// карта частот
	stats.wordFreq = make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stats.LineCount++
		stats.CharCount += len([]rune(line))

		// строку на слова
		words := strings.Fields(line)
		for _, word := range words {
			// нижний регистр и убираем знаки пунктуации
			word = strings.ToLower(strings.Trim(word, ".,!?;:\"'()[]{}«»—–-"))
			if word == "" {
				continue
			}
			stats.WordCount++
			stats.wordFreq[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		return stats, err
	}

	return stats, nil
}

func printStats(stats TextStats) {
	fmt.Println("__  Статистика текста __")
	fmt.Printf("Символов: %d\n", stats.CharCount)
	fmt.Printf("Слов:     %d\n", stats.WordCount)
	fmt.Printf("Строк:    %d\n", stats.LineCount)
}

// getTopWords возвращает topN самых частых слов
func getTopWords(wordFreq map[string]int, topN int) []WordStat {
	result := make([]WordStat, 0, len(wordFreq))

	// переложить карту в срез
	for word, count := range wordFreq {
		result = append(result, WordStat{Word: word, Count: count})
	}

	// сортировка по убыванию Count
	sort.Slice(result, func(i, j int) bool {
		return result[i].Count > result[j].Count
	})

	// вернуть topN
	if topN > len(result) {
		topN = len(result)
	}
	return result[:topN]
}

// printTopWords печатает топ слов
func printTopWords(top []WordStat) {
	fmt.Println("__ Топ-5 слов __")
	for i, ws := range top {
		fmt.Printf("%d. %s: %d\n", i+1, ws.Word, ws.Count)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование: go-text <файл>")
		os.Exit(1)
	}

	stats, err := analyzeFile(os.Args[1])
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		os.Exit(1)
	}

	printStats(stats)

	top := getTopWords(stats.wordFreq, 5)
	printTopWords(top)
}