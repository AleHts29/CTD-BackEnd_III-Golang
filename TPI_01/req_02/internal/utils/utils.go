package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseTime(timeStr string) (hour, minute int, err error) {
	parts := strings.Split(timeStr, ":")
	if len(parts) != 2 {
		err = fmt.Errorf("formato de tiempo incorrecto: %s", timeStr)
		return
	}
	hour, err = strconv.Atoi(parts[0])
	//hour, err = time.Parse("15", parts[0])
	if err != nil {
		return
	}
	minute, err = strconv.Atoi(parts[1])
	return
}
