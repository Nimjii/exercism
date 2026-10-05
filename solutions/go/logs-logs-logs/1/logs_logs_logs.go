package logs

import (
    "fmt"
    "strings"
    "unicode/utf8"
)
    
func Application(log string) string {
	for _, char := range(log) {
        switch fmt.Sprintf("%U", char) {
        case "U+2757":
            return "recommendation"
        case "U+1F50D":
            return "search"
        case "U+2600":
            return "weather"
        }
    }
    return "default"
}

func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
