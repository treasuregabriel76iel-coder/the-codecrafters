package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func upper(words string) string {
	return strings.ToUpper(words)
}
func lower(words string) string {
	return strings.ToLower(words)
}
func cap(words string) string {
	text := strings.Fields(words)
	for i, item := range text {
		text[i] = strings.ToUpper(string(item[0])) + strings.ToLower(item[1:])
	}
	return strings.Join(text, " ")
}
func title(words string) string {
	text := strings.Fields(words)
	small := map[string]bool{
		"a": true, "an": true, "the": true,
		"and": true, "but": true, "or": true,
		"for": true, "nor": true,
		"on": true, "at": true, "to": true,
		"by": true, "in": true, "of": true,
		"up": true, "as": true, "is": true, "it": true,
	}
	for i, w := range text {
		line := strings.ToLower(w)
		if i == 0 {
			text[i] = strings.ToUpper(string(line[0])) + line[1:]
			continue
		}
		if small[line] {
			text[i] = line
		} else {
			text[i] = strings.ToUpper(string(line[0])) + line[1:]
		}
	}
	return strings.Join(text, " ")
}

// func snake(words string) string {
// }
func reverse(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		runes := []rune(w)
		for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
			runes[l], runes[r] = runes[r], runes[l]
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("SENTINEL STRING TRANSFORMER — ONLINE")
	fmt.Println("Supported commands: ")
	fmt.Println("(up) (text) - uppercase")
	fmt.Println("(low) (text) - lowercase")
	fmt.Println("(cap) (text) - capitalize first letter of each word in a string")
	fmt.Println("(title) (text) -  ")
	fmt.Println("(snake) (text) - snakecase")
	fmt.Println("(rev) (text) - reverse")
	fmt.Println("type 'exit' to quit :")
	for {
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if strings.ToLower(line) == "exit" {
			fmt.Println("Shutting down String Transformer. Goodbye.")
			break
		}
		parts := strings.Fields(line)
		command := strings.ToLower(parts[0])
		if len(parts) == 1 {
			fmt.Printf("No text provided. Usage: %s <text>\n", command)
			continue
		}
		text := strings.Join(parts[1:], " ")
		switch command {
		case "up":
			fmt.Println(upper(text))
		case "low":
			fmt.Println(lower(text))
		case "cap":
			fmt.Println(cap(text))
		case "title":
			fmt.Println(title(text))
		// case "snake":
		// 	fmt.Println(snake(text))
		case "rev":
			fmt.Println(reverse(text))
		default:
			fmt.Printf("Unknown command: %s\n", command)
			fmt.Println("Valid commands: up, low, cap, title, snake, rev, exit")
		}
	}
}
