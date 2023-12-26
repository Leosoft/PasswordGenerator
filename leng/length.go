package leng

import (
	"fmt"
	"os"
	"strconv"
)

func GetLength() (int, error) {
	if len(os.Args) < 2 {
		return 0, fmt.Errorf("Please check password length and retry")
	}

	length, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return 0, fmt.Errorf("length value must be integer")
	}

	return length, nil
}
