package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

type cannelName struct {
	Name string
	Channel string
}

func (s cannelName) String() string {
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
	var name string
	fmt.Scanf("%s\n", &name)

	file, err := os.Open("russian_nouns.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	// Создаем сканер для построчного чтения
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if channel, ok := findChannel(line, name); ok {
			fmt.Println(channel)
		}
	}


	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при сканировании файла:", err)
	}
}

func findChannel(line, name string) (cannelName, bool) {
	if (len(line) <= len(name)) {
		return cannelName{}, false
	}
	j := 0
	for i := 0; i < len(line); i++ {
		if j < len(name) && line[i] == name[j] {
			j++
		}
	}

	if (j == len(name)) {
		return cannelName{Name: name, Channel: line}, true
	} 
	return cannelName{}, false
}