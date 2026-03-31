    // CodeCrafters — Operation Gopher Protocol
       // Module: File Pipeline
       // Author: [Treasure Gabriel]
       // Squad:  [Benchmarkers]


package main
import (
	"fmt"
	"bufio"
	"strings"
	"os"
)
 // ═══════════════════════════════════════════
  // SQUAD PIPELINE CONTRACT
  // Squad: [Squad Name]
  // ───────────────────────────────────────────
  // Input line types:
  //   [line 1: ALL CAP]
  //   [line 2: lower]
  //   [line 3: Trimspace]
  //   [line 4: TODO ]
  //
  // Transformation rules (in order):
  //   1. [Trim all leading and trailing whitespace ]
  //   2. [Replace TODO: with ✦ ACTION]
  //   3. [Convert ALL CAPS lines to Title Case ]
  //   4. [Convert all lowercase lines to uppercase ]
  //   5. [Remove lines that are only dashes or blanks 	]
  //
  // Output format:
  //   [Header: yes/no — SENTINEL FIELD REPORT — PROCESSED]
  //   [Line numbering format- 1]
  //
  // Terminal summary fields:
  //    ✦ Lines read    : [number]                  						
   //   ✦ Lines written : [number]                  						
   //  ✦ Lines removed : [number]                  					
  //   ✦ Rules applied :   
  //   1. [Trim all leading and trailing whitespace ]
  //   2. [Replace TODO: with ✦ ACTION]
  //   3. [Convert ALL CAPS lines to Title Case ]
  //   4. [Convert all lowercase lines to uppercase ]
  //   5. [Remove lines that are only dashes or blanks 	]
  // ═══════════════════════════════════════════

func Trimspace(s string) string {
	return strings.TrimSpace(s)
}

func replaceTODO(s string) string {
	if strings.HasPrefix(s, "TODO:") {
		return strings.Replace(s, "TODO:", "✦ ACTION:", 1)
	}
	return s
}

func capToTitle(s string) string {
	if s == strings.ToUpper(s) && s != "" {
		return Title(s)
	}
	return s
}

func lowerToUpper(s string) string {
	if s == strings.ToLower(s) && s != "" {
		return strings.ToUpper(s)
	}
	return s
}

func lineRemove(s string) bool {
	if s == "" {
		return true
	}

	onlyDashes := true
	for _, ch := range s {
		if ch != '-' {
			onlyDashes = false
			break
		}
	}
	return onlyDashes
}

func Title(s string) string {
	words := strings.Fields(strings.ToLower(s))
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(string(w[0])) + w[1:]
		}
	}
	return strings.Join(words, " ")
}
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if inputFile == outputFile {
		fmt.Println("✗ Input and output cannot be the same file.")
		return
	}

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("✗ File not found: %s\n", inputFile)
		return
	}
	defer file.Close()

	info, err := os.Stat(outputFile)
	if err == nil && info.IsDir() {
		fmt.Println("✗ Cannot write to output: path is a directory.")
		return
	}

	out, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("✗ Error creating output file.")
		return
	}
	defer out.Close()

	writer := bufio.NewWriter(out)
	defer writer.Flush()

	scanner := bufio.NewScanner(file)

	linesRead := 0
	linesWritten := 0
	linesRemoved := 0

	var processed []string

	for scanner.Scan() {
		line := scanner.Text()
		linesRead++

		line = Trimspace(line)
		line = replaceTODO(line)
		line = capToTitle(line)
		line = lowerToUpper(line)

		if lineRemove(line) {
			linesRemoved++
			continue
		}

		processed = append(processed, line)
	}

	if linesRead == 0 {
		fmt.Println("⚠ Input file is empty. Nothing to process.")
	}

	writer.WriteString("SENTINEL FIELD REPORT — PROCESSED\n\n")

	for i, line := range processed {
		numbered := fmt.Sprintf("%d. %s\n", i+1, line)
		writer.WriteString(numbered)
		linesWritten++
	}

	fmt.Printf("✦ Lines read    : %d\n", linesRead)
	fmt.Printf("✦ Lines written : %d\n", linesWritten)
	fmt.Printf("✦ Lines removed : %d\n", linesRemoved)
	fmt.Println("✦ Rules applied :")
	fmt.Println("  1. Trim all leading and trailing whitespace")
	fmt.Println("  2. Replace \"TODO:\" with \"✦ ACTION:\"")
	fmt.Println("  3. Convert ALL CAPS lines to Title Case")
	fmt.Println("  4. Convert all lowercase lines to UPPERCASE")
	fmt.Println("  5. Remove lines that are only dashes or blank")
}
