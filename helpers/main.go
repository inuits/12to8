package helpers

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// GetMonthYearFromArg returns the year and the month represented by the user-defined arg
func GetMonthYearFromArg(arg string) (int, int, error) {
	if arg == "" {
		return int(time.Now().Month()), time.Now().Year(), nil
	}
	v := strings.Split(arg, "/")
	if len(v) > 2 {
		return 0, 0, fmt.Errorf("Too many / in %v", v)
	}
	var y int
	var err error
	if len(v) == 1 {
		y = time.Now().Year()
	} else {
		y, err = strconv.Atoi(v[1])
		if err != nil {
			return 0, 0, err
		}
	}
	m, err := strconv.Atoi(v[0])
	if err != nil {
		return 0, 0, err
	}
	if m > 12 || m < 1 {
		return 0, 0, fmt.Errorf("bad month: %v", v[0])
	}
	return m, y, nil
}
