package services

import (
	"fmt"
	"strings"
	"time"
)

func GenerateFileName(fileName string) string {
	fileName = strings.ReplaceAll(fileName, ".md", "")
	now := time.Now()
	timestamp := now.UnixMicro()

	// Returns the template name with a timestamp
	return fmt.Sprintf("%s %d", fileName, timestamp)
}

// Temporary solution to pass a filename as an arg (without spaces)
func GetActionableFileName(fileName string) string {
	return strings.ReplaceAll(fileName, " ", "")
}
