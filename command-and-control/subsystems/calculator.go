package subsystems

import "errors"

func Calc(op string, a, b float64) (float64, error) {
	switch op {
	case "add":
		return a + b, nil
	case "sub":
		return a - b, nil
	case "mul":
		return a * b, nil
	case "div":
		if b == 0 {
			return 0, errors.New("cannot divide by zero")
		}
		return a / b, nil
	case "mod":
		if b == 0 {
			return 0, errors.New("cannot mod by zero")
		}
		return float64(int(a) % int(b)), nil
	case "pow":
		result := 1.0
		for i := 0; i < int(b); i++ {
			result *= a
		}
		return result, nil
	default:
		return 0, errors.New("unknown operation")
	}
}