package stringtransformer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func upper(s string) string {
	return strings.ToUpper(s)
}
func lower(s string) string {
	return strings.ToLower(s)
}
func cap(s string) string {
	text := strings.Fields(s)
	for i, item := range text {
		text[i] = strings.ToUpper(string(item[0])) + strings.ToLower(item[1:])
	}
	return strings.Join(text, " ")
}
func title(words string) string {

	result := strings.Fields(words)

	for i := 0; i < len(result); i++ {
		if len(result[i]) < 3 && i != 0 {
			result[i] = strings.ToLower(result[i])
		} else {
			result[i] = cap(result[i])
		}
	}
	return strings.Join(result, " ")
}
func snake(words string) string {

}
func reverse(words string) string {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if strings.ToLower(line) == "exit" {
			fmt.Println("GOODBYE!!")
			break
		}
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Error: Invalid format. Usage: <number> <base>")
			continue
		}

	}
}
