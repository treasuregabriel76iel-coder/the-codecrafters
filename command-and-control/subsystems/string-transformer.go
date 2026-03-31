package subsystems

import (
	"errors"
	"strconv"
	"strings"
)

func Transform(cmd string, text string, last float64) (string, error) {
	if strings.Contains(text, "last") {
		text = strings.ReplaceAll(text, "last", strconv.Itoa(int(last)))
	}

	switch cmd {
	case "upper":
		return strings.ToUpper(text), nil
	case "lower":
		return strings.ToLower(text), nil
	case "cap":
		words := strings.Fields(text)
		for i, w := range words {
			words[i] = strings.ToUpper(string(w[0])) + strings.ToLower(w[1:])
		}
		return strings.Join(words, " "), nil
	case "title":
		words := ""
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
				text[i] = strings.ToUpper(string(line[0])) + strings.ToLower(line[1:])
			}
			return strings.Join(text, " "), nil
		}

	case "snake":
		return strings.ToLower(strings.ReplaceAll(text, " ", "_")), nil
	case "reverse":
		words := strings.Fields(text)
		for i, w := range words {
			r := []rune(w)
			for l, h := 0, len(r)-1; l < h; l, h = l+1, h-1 {
				r[l], r[h] = r[h], r[l]
			}
			words[i] = string(r)
		}
		return strings.Join(words, " "), nil
	default:
		return "", errors.New("unknown string command")
	}
	return "", nil
}