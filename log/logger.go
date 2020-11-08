package log

import (
	"fmt"
	"os"
)

func Die(msg string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, msg, v)
	os.Exit(1)
}

func Exit(msg string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, msg, v)
	os.Exit(0)
}

func printLevelLog(level string, msg string, v ...interface{}) {
	m := fmt.Sprintf(msg, v)
	fmt.Fprintln(os.Stderr, "%s: %s", level, m)
}

func Debug(msg string, v ...interface{}) {
	printLevelLog("DEBUG", msg, v)
}

func Warning(msg string, v ...interface{}) {
	printLevelLog("WARNING", msg, v)
}
