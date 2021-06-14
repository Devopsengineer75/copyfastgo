package lib

import "errors"

func division(a int, b int) (error, int) {

	if b == 0 {
		return errors.New("Division par 0 impossible"), 0
	}

	return nil, a / b
}
