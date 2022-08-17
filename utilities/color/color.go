package color

import (
	"fmt"
	"io"
	"runtime"

	"github.com/goEnum/goEnum/utilities/parameters"
)

type colorString string

const (
	Reset                  colorString = "\033[0m"
	Black                  colorString = "\033[31m"
	Red                    colorString = "\033[31m"
	Green                  colorString = "\033[32m"
	Yellow                 colorString = "\033[33m"
	Blue                   colorString = "\033[34m"
	Purple                 colorString = "\033[35m"
	Cyan                   colorString = "\033[36m"
	Gray                   colorString = "\033[37m"
	White                  colorString = "\033[97m"
	Bold                   colorString = "\033[1m"
	Dim                    colorString = "\033[2m"
	Italic                 colorString = "\033[3m"
	Underline              colorString = "\033[4m"
	SlowBlink              colorString = "\033[5m"
	FastBlink              colorString = "\033[6m"
	Invisible              colorString = "\033[8m"
	Stikeout               colorString = "\033[9m"
	DoubleUnderline        colorString = "\033[0m"
	BlackBackground        colorString = "\033[40m"
	RedBackground          colorString = "\033[41m"
	GreenBackground        colorString = "\033[42m"
	YellowBackground       colorString = "\033[43m"
	BlueBackground         colorString = "\033[44m"
	PurpleBackground       colorString = "\033[45m"
	CyanBackground         colorString = "\033[46m"
	WhiteBackground        colorString = "\033[47m"
	BrightBlack            colorString = "\033[91m"
	BrightRed              colorString = "\033[91m"
	BrightGreen            colorString = "\033[92m"
	BrightYellow           colorString = "\033[93m"
	BrightBlue             colorString = "\033[94m"
	BrightPurple           colorString = "\033[95m"
	BrightCyan             colorString = "\033[96m"
	BrightBlackBackground  colorString = "\033[100m"
	BrightRedBackground    colorString = "\033[101m"
	BrightGreenBackground  colorString = "\033[102m"
	BrightYellowBackground colorString = "\033[103m"
	BrightBlueBackground   colorString = "\033[104m"
	BrightPurpleBackground colorString = "\033[105m"
	BrightCyanBackground   colorString = "\033[106m"
	BrightWhiteBackground  colorString = "\033[107m"
)

func Fprintln(color colorString, w io.Writer, a ...any) (n int, err error) {
	if parameters.Color && runtime.GOOS != "windows" {
		return fmt.Fprintf(w, "%v%v%v", color, fmt.Sprintln(a...), Reset)
	} else {
		return fmt.Fprintln(w, a...)
	}
}

func Println(color colorString, a ...any) (n int, err error) {
	if parameters.Color && runtime.GOOS != "windows" {
		return fmt.Printf("%v%v%v\n", color, fmt.Sprint(a...), Reset)
	} else {
		return fmt.Println(a...)
	}
}

func Fprintf(color colorString, w io.Writer, format string, a ...any) (n int, err error) {
	if parameters.Color && runtime.GOOS != "windows" {
		return fmt.Fprintf(w, "%v%v%v", color, fmt.Sprintf(format, a...), Reset)
	} else {
		return fmt.Fprintf(w, format, a...)
	}
}

func Printf(color colorString, format string, a ...any) (n int, err error) {
	if parameters.Color && runtime.GOOS != "windows" {
		return fmt.Printf("%v%v%v", color, fmt.Sprintf(format, a...), Reset)
	} else {
		return fmt.Printf(format, a...)
	}
}

func Sprintf(color colorString, format string, a ...any) string {
	if parameters.Color && runtime.GOOS != "windows" {
		return fmt.Sprintf("%v%v%v", color, fmt.Sprintf(format, a...), Reset)
	} else {
		return fmt.Sprintf(format, a...)
	}

}
func Sprintln(color colorString, a ...any) string {
	if parameters.Color && runtime.GOOS != "windows" {
		return fmt.Sprintf("%v%v%v\n", color, fmt.Sprint(a...), Reset)
	} else {
		return fmt.Sprintln(a...)
	}
}
