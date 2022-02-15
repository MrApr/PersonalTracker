package Error

import (
	"fmt"
	"github.com/MrApr/PersonalTracker/logger"
	"runtime/debug"
)

//AdvanceError Defines Custom error type
type AdvanceError struct {
	Message string
	File    string
	Line    int
	Type    string
}

//AdvanceError implements error interface for error struct implicitly
func (e *AdvanceError) Error() string {
	logger := logger.Create("default")
	var format string
	format = "%s: %s, %s in line %d\n"
	logger.Printf(format, e.Type, e.File, e.Message, e.Line)
	debug.Stack()
	return fmt.Sprintf(format, e.Type, e.File, e.Message, e.Line)
}
