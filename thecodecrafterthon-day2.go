package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func enter() {
	const (
		Red    = "\033[0;31m"
		Green  = "\033[0;32m"
		Yellow = "\033[0;33m"
		Blue   = "\033[0;34m"
		Reset  = "\033[0m"
		Purple = "\033[35m"
	)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(Purple + "***** WELCOME TO MY CLI BASE CONVERTER *****" + Reset)
	fmt.Println(Blue + "✦✦✦ USAGE :(number) (base) ✦✦✦" + Reset)
	fmt.Println(Purple + "Supported commands:" + Reset)
	fmt.Println(Green + "  (number) (hex)  - hexadecimal" + Reset)
	fmt.Println(Yellow + "  (number) (bin) - binary" + Reset)
	fmt.Println(Purple + "  (number) (dec) - decimal " + Reset)
	fmt.Println(Yellow + "type 'exit' to quit :" + Reset)
	for {
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if strings.ToLower(line) == "exit" {
			fmt.Println(Purple + "GOODBYE!!")
			break
		}
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println(Red + "Error: Invalid format. Usage: <number> <base>" + Reset)
			continue
		}
		numStr := parts[0]
		base := strings.ToLower(parts[1])
		baseConversions(numStr, base)
	}
}
func baseConversions(numstr, base string) {
	const (
		Red    = "\033[0;31m"
		Green  = "\033[0;32m"
		Yellow = "\033[0;33m"
		Blue   = "\033[0;34m"
		Reset  = "\033[0m"
		Purple = "\033[35m"
	)
	var decimalvalue int64
	var err error
	switch base {
	case "dec":
		decimalvalue, err = strconv.ParseInt(numstr, 10, 64)
	case "hex":
		decimalvalue, err = strconv.ParseInt(numstr, 16, 64)
	case "bin":
		decimalvalue, err = strconv.ParseInt(numstr, 2, 64)
	default:
		fmt.Println(Red+"Error unsupported base. Available base : dec, hex, bin\n", base+Reset)
		return
	}
	if err != nil {
		fmt.Printf(Red+"Error: '%s' is not valid %s number.\n", numstr, base+Reset)
	}
	switch base {
	case "dec":
		binary := strconv.FormatInt(decimalvalue, 2)
		hex := strings.ToUpper(strconv.FormatInt(decimalvalue, 16))
		fmt.Printf("Binary:  %s\n", binary)
		fmt.Printf("Hex:     %s\n", hex)
	case "hex", "bin":
		fmt.Printf("Decimal: %d\n", decimalvalue)
	}
}
