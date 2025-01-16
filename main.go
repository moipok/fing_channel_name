package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type channelnfo struct {
	Name string
	Channel string
}

func (s channelnfo) String() string {
	var builder strings.Builder
	nameRunes := []rune(s.Name)
	channelRunes := []rune(s.Channel)
	j := 0

	for i := 0; i < len(channelRunes); i++ {
		if j < len(nameRunes) && channelRunes[i] == nameRunes[j] {
			builder.WriteString(strings.ToUpper(string(nameRunes[j])))
			j++
		} else {
			builder.WriteString(string(channelRunes[i]))
		}
	}

	return builder.String()
}

func main() {
	name := readUserName()
	file := openFile("russian_nouns.txt")
	defer file.Close()

	results := processFile(file, name)
	printResults(results)
}

func readUserName() string {
	var name string
	fmt.Print("Введите имя: ")
	fmt.Scanf("%s\n", &name)
	return strings.ToLower(name)
}

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		os.Exit(1)
	}
	return file
}

func processFile(file *os.File, name string) chan channelnfo {
	results := make(chan channelnfo)
	var wg sync.WaitGroup

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			if channel, ok := findChannel(line, name); ok {
				results <- channel
			}
		}(line)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при сканировании файла:", err)
	}

	return results
}

func printResults(results chan channelnfo) {
	count := 0
	for result := range results {
		count++
		fmt.Println(count, "\t", result)
	}
}

func findChannel(line, name string) (channelnfo, bool) {
	if (len(line) <= len(name)) {
		return channelnfo{}, false
	}
	j := 0
	for i := 0; i < len(line); i++ {
		if j < len(name) && line[i] == name[j] {
			j++
		}
	}

	if (j == len(name)) {
		return channelnfo{Name: name, Channel: line}, true
	} 
	return channelnfo{}, false
}