package logger

import (
	"fmt"
	"os"
	"strconv"
)

func Log(body string, colorCode int) {
	fmt.Println("\u001b[" + strconv.Itoa(colorCode) + "m" + body + "\u001b[0m")
}

func Err(body string) {
	Log(body, 31)
	os.Exit(0)
}

func Ok(body string) {
	Log(body, 32)
}

func Info(body string) {
	Log(body, 34)
}
