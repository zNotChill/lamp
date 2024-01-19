package utils

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func Info(message string) {
	lime := color.New(color.FgHiGreen).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()

	t := time.Now()
	fmt.Printf("[%s] [%s] %s\n", lime(t.Format("15:04:05")), blue("INFO"), white(message))
}

func Warn(message string) {
	lime := color.New(color.FgHiGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()

	t := time.Now()
	fmt.Printf("[%s] [%s] %s\n", lime(t.Format("15:04:05")), yellow("WARN"), white(message))
}

func Error(message string) {
	lime := color.New(color.FgHiGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()

	t := time.Now()
	fmt.Printf("[%s] [%s] %s\n", lime(t.Format("15:04:05")), red("ERROR"), white(message))
}

func Fatal(message string) {
	lime := color.New(color.FgHiGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()

	t := time.Now()
	fmt.Printf("[%s] [%s] %s\n", lime(t.Format("15:04:05")), red("FATAL"), white(message))
}
