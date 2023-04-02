package log

import (
	"fmt"
	"runtime"
	"time"

	"github.com/fatih/color"
)

func Log(message string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)

	if ok {
		fmt.Println(color.GreenString("%s:%d", file, line))
	}

	_, file, line, ok = runtime.Caller(2)

	if ok {
		fmt.Println(color.GreenString("%s:%d", file, line))
	}

	_, file, line, ok = runtime.Caller(3)

	if ok {
		fmt.Println(color.GreenString("%s:%d", file, line))
	}

	_, file, line, ok = runtime.Caller(4)

	if ok {
		fmt.Println(color.GreenString("%s:%d", file, line))
	}

	if message[0] == '#' {
		if message[1] == 'f' {
			message = message[2:]
			message = color.RedString("🔥%s", message)
		}

		if message[1] == 'e' {
			message = message[2:]
			message = color.GreenString("👁%s", message)
		}

		if message[1] == 'd' {
			message = message[2:]
			message = color.BlueString("🔵%s", message)
		}

		if message[1] == 'i' {
			message = message[2:]
			message = color.YellowString("💧%s", message)
		}
	}

	fmt.Print(color.BlueString(time.Now().Format("[15:04:05.000000] ")))
	fmt.Println(fmt.Sprintf(message, args...))
}

func Time(label string) {

	fmt.Println(fmt.Sprintf("%s: %s", color.YellowString(time.Now().Format("[15:04:05.000000] ")), label))
}
