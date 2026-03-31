// CodeCrafters — Hackathon 002
// Squad: [BenchMarkers]
// Members: [ Peace Oigoga, Faith Ejembi, Nzekwe Chinaza, Treasure Gabriel, Owulo Celebrate, Ortse Alphonsus, Chekwube Okafor, Emaikwu Godwin, Endurance Ochefije, Chubiyojo Akoh]

package main

import (
	"bufio"
	"fmt"
	"os"
	"sentinel-control/subsystems"
	"strconv"
	"strings"
)

var last float64
var history []string

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("  SENTINEL — COMMAND & CONTROL CONSOLE")
	fmt.Println("     All systems nominal. Type 'help' to begin.")
	for {

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}
		addHistory(input)

		switch input {
		case "exit":
			fmt.Println("Leaving so soon?...The Benchmarkers hope to see you soon. GOODBYE!!😟")
			return

		case "help":
			printHelp()
			continue

		case "history":
			printHistory()
			continue
		case "clear":
			handleClear(scanner)
			continue
		}
		if strings.Contains(input, "|") {
			handlePipe(input)
			continue
		}
		execute(input, "")
	}
}

func execute(input string, piped string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return
	}
	module := strings.ToLower(parts[0])
	switch module {
	case "calc":
		if len(parts) < 2 {
			fmt.Println("✗ Error: missing calc command")
			return
		}
		cmd := parts[1]
		if cmd == "last" {
			fmt.Println("✦ Result:", last)
			return
		}
		if len(parts) < 4 {
			fmt.Println("✗ Error: missing arguments")
			return
		}
		a, err1 := resolveNumber(parts[2], piped)
		b, err2 := resolveNumber(parts[3], piped)

		if err1 != nil || err2 != nil {
			fmt.Println("✗ Error: invalid number")
			return
		}
		result, err := subsystems.Calc(cmd, a, b)
		if err != nil {
			fmt.Println("✗ Error:", err)
			return
		}
		last = result
		fmt.Println("✦ Result:", result)

	case "base":
		if len(parts) < 2 {
			fmt.Println("✗ Error: missing base command")
			return
		}

		cmd := parts[1]

		var value string

		if len(parts) >= 3 {
			value = parts[2]
		} else {
			value = piped
		}

		if value == "" {
			fmt.Println("✗ Error: missing number")
			return
		}

		output, dec, err := subsystems.BaseConvert(cmd, value)
		if err != nil {
			fmt.Println("✗ Error:", err)
			return
		}
		last = float64(dec)
		fmt.Print(output)
	case "str":
		if len(parts) < 2 {
			fmt.Println("✗ Error: missing string command")
			return
		}

		cmd := parts[1]

		text := ""
		if len(parts) > 2 {
			text = strings.Join(parts[2:], " ")
		} else {
			text = piped
		}

		if text == "" {
			fmt.Println("✗ Error: no text provided")
			return
		}

		result, err := subsystems.Transform(cmd, text, last)
		if err != nil {
			fmt.Println("✗ Error:", err)
			return
		}

		fmt.Println("✦", result)

	default:
		fmt.Println("✗ Unknown command. Type 'help'")
	}
}

func handlePipe(input string) {
	parts := strings.Split(input, "|")
	left := strings.TrimSpace(parts[0])
	right := strings.TrimSpace(parts[1])
	oldLast := last
	execute(left, "")
	piped := fmt.Sprintf("%.0f", last)
	execute(right, piped)
	last = last + oldLast
}

func resolveNumber(val string, piped string) (float64, error) {
	if val == "last" {
		return last, nil
	}
	if val == "" {
		val = piped
	}
	return strconv.ParseFloat(val, 64)
}

func addHistory(cmd string) {
	if len(history) >= 10 {
		history = history[1:]
	}
	history = append(history, cmd)
}

func printHistory() {
	for i, h := range history {
		fmt.Printf("%d. %s\n", i+1, h)
	}
}

func handleClear(scanner *bufio.Scanner) {
	fmt.Print("Type CONFIRM to clear the session: ")
	if !scanner.Scan() {
		return
	}
	if scanner.Text() == "CONFIRM" {
		history = nil
		last = 0
		fmt.Println("Session cleared.")
	} else {
		fmt.Println("Cancelled.")
	}
}
func printHelp() {
	fmt.Println("calc add/sub/mul/div/mod/pow")
	fmt.Println("base dec/hex/bin")
	fmt.Println("str upper/lower/cap/title/snake/reverse")
	fmt.Println("history, clear, exit")
}