package utilities

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/goEnum/goEnum/structs"
)

func exists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func Append(path string, text bytes.Buffer) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer file.Close()

	io.Copy(file, &text)
}

func Write(path string, text bytes.Buffer) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	defer file.Close()

	io.Copy(file, &text)
}

func Read(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return string(data)
}

func WriteJSON(path string, report *structs.JSONReport) {
	var (
		reports []structs.JSONReport
		initial string
	)

	if exists(path) {
		initial = Read(path)

		err := json.Unmarshal([]byte(initial), &reports)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	reports = append(reports, *report)

	data, err := json.Marshal(reports)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	var buffer bytes.Buffer
	fmt.Fprint(&buffer, data)

	Write(path, buffer)
}
