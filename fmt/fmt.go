package fmt

import (
	"fmt"
	"github.com/fatih/color"
)

var (
	yellow func(a ...interface{}) string = color.New(color.FgYellow).SprintFunc()
	red    func(a ...interface{}) string = color.New(color.FgRed).SprintFunc()
	green  func(a ...interface{}) string = color.New(color.FgGreen).SprintFunc()
)

type Logger interface {
	Error(msg string) string
	Complete(msg string) string
	Incomplete(msg string) string
}

type SimpleLogger struct{}
type AsciiLogger struct{}
type PrettyLogger struct{}

func NewLogger(kind string) Logger {
	if kind == "" {
		return nil
	}

	var f Logger

	if kind == "simple" {
		f = &SimpleLogger{}
	} else if kind == "ascii" {
		f = &AsciiLogger{}
	} else if kind == "pretty" {
		f = &PrettyLogger{}
	} else {
		panic("Invalid logger passed")
	}

	return f
}

//
// SimpleLogger methods
//
func (l *SimpleLogger) Error(msg string) string {
	return fmt.Sprintf("Error: %s", msg)
}

func (l *SimpleLogger) Complete(msg string) string {
	return fmt.Sprintf("Complete: %s", msg)
}

func (l *SimpleLogger) Incomplete(msg string) string {
	return fmt.Sprintf("Incomplete: %s", msg)
}

//
// PrettyLogger methods
//
func (l *PrettyLogger) Error(msg string) string {
	return fmt.Sprintf("✖ Error: %s", msg)
}

func (l *PrettyLogger) Complete(msg string) string {
	return fmt.Sprintf("✓ Complete: %s", msg)
}

func (l *PrettyLogger) Incomplete(msg string) string {
	return fmt.Sprintf("✖ Incomplete: %s", msg)
}

//
// AsciiLogger methods
//
func (l *AsciiLogger) Error(msg string) string {
	return fmt.Sprintf("%s %s", red("✖ Error:"), msg)
}

func (l *AsciiLogger) Complete(msg string) string {
	return fmt.Sprintf("%s %s", green("✓ Complete:"), msg)
}

func (l *AsciiLogger) Incomplete(msg string) string {
	return fmt.Sprintf("%s %s", yellow("✖ Incomplete:"), msg)
}
