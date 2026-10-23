package main

import (
	"fmt"
	"os"
	"strings"
)

func loadBanner(bannerFile string) (map[rune][]string, error) {
	data, err := os.ReadFile(bannerFile)
	if err != nil {
		return nil, err
	}

	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	content = strings.TrimRight(content, "\n")

	linesFile := strings.Split(content, "\n")

	alphabet := make(map[rune][]string)
	charCode := 32
	lineIndex := 1

	for charCode <= 126 && lineIndex+8 <= len(linesFile) {
		charLines := linesFile[lineIndex : lineIndex+8]
		alphabet[rune(charCode)] = charLines

		lineIndex += 9
		charCode++
	}

	return alphabet, nil
}

func renderASCII(input string, alphabet map[rune][]string) (string, error) {
	var sb strings.Builder

	lines := strings.Split(input, "\n")

	for i, line := range lines {
		if line == "" {
			if i != 0 {
				sb.WriteString("\n")
			}
			continue
		}

		for row := 0; row < 8; row++ {
			for _, ch := range line {
				linesForChar, ok := alphabet[ch]
				if !ok {
					return "", fmt.Errorf("unsupported character: %q", ch)
				}
				sb.WriteString(linesForChar[row])
			}
			sb.WriteString("\n")
		}
	}

	return sb.String(), nil
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
		return
	}

	bannerFile := "standard.txt"
	if len(os.Args) == 3 {
		switch os.Args[2] {
		case "standard":
			bannerFile = "standard.txt"
		case "shadow":
			bannerFile = "shadow.txt"
		case "thinkertoy":
			bannerFile = "thinkertoy.txt"
		}
	}

	input := os.Args[1]
	input = strings.ReplaceAll(input, `\n`, "\n")

	alphabet, err := loadBanner(bannerFile)
	if err != nil {
		return
	}

	output, err := renderASCII(input, alphabet)
	if err != nil {
		return
	}

	fmt.Print(output)
}
