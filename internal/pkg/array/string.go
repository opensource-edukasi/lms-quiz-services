package array

import (
	"fmt"
	"strings"
)

func RemoveByValue(arr []string, val string) []string {
	for i, v := range arr {
		if v == val {
			return append(arr[:i], arr[(i+1):]...)
		}
	}

	return arr
}

func ConvertToWhereIn(i int, arr []string) string {
	var rep []string
	for _, _ = range arr {
		rep = append(rep, fmt.Sprintf("$%d", i))
		i++
	}
	return strings.Join(rep, ",")
}

func ConvertToAny(arr []string) []any {
	var rep []any
	for _, v := range arr {
		rep = append(rep, v)
	}
	return rep
}
