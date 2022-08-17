package utilities

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Contains(s []string, str string) bool {
	if str == "" {
		return true
	}

	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func ListPrint(title string, list []string) string {
	var output bytes.Buffer

	if len(list) == 0 {
		fmt.Fprintf(&output, "%v:\n", title)
		return output.String()
	}

	fmt.Fprintf(&output, "%v: %v\n", title, list[0])

	padding := strings.Repeat(" ", len(title))
	padding += "- "

	for _, element := range list[1:] {
		fmt.Fprintf(&output, "%v%v\n", padding, element)
	}

	return output.String()
}

func RandString(length int) string {
	rand.Seed(time.Now().UnixNano())

	runes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456790")
	output := make([]rune, length)

	for index := range output {
		output[index] = runes[rand.Intn(len(runes))]
	}

	return string(output)
}
