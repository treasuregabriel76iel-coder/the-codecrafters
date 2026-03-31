package subsystems

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func BaseConvert(mode string, input string) (string, int64, error) {
	var dec int64
	var err error

	switch mode {
	case "dec":
		dec, err = strconv.ParseInt(input, 10, 64)
		if err != nil {
			return "", 0, errors.New("invalid decimal")
		}
		bin := strconv.FormatInt(dec, 2)
		hex := strings.ToUpper(strconv.FormatInt(dec, 16))
		return fmt.Sprintf("✦ Binary : %s\n✦ Hex    : %s\n", bin, hex), dec, nil

	case "hex":
		dec, err = strconv.ParseInt(input, 16, 64)
		if err != nil {
			return "", 0, errors.New("invalid hex")
		}
		return fmt.Sprintf("✦ Decimal: %d\n", dec), dec, nil

	case "bin":
		dec, err = strconv.ParseInt(input, 2, 64)
		if err != nil {
			return "", 0, errors.New("invalid binary")
		}
		return fmt.Sprintf("✦ Decimal: %d\n", dec), dec, nil

	default:
		return "", 0, errors.New("unknown base command")
	}
}