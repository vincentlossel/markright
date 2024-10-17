package services

import (
	"fmt"
	"strings"
	"time"
)

func GenerateFileName(fileName string, addTimestamp bool) string {
	fileName = strings.ReplaceAll(fileName, ".md", "")

	// Timestamp is optional
	if addTimestamp {
		now := time.Now()
		timestamp := now.UnixMicro()

		newFileName := fmt.Sprintf("%s %d", fileName, timestamp)
		return newFileName
	}

	return fmt.Sprintf("%s", fileName)
}

// Temporary solution to pass a filename as an arg (without spaces)
func GetActionableFileName(fileName string) string {
	fileName = strings.ReplaceAll(fileName, " ", "")
	fileName = strings.ReplaceAll(fileName, ".md", "")

	return fileName
}
