package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
    for _, runeValue := range log {
        if runeValue == '❗' {
            return "recommendation"
        } else if runeValue == '🔍' {
            return "search"
        } else if runeValue == '☀' {
            return "weather"
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    foo := make([]rune, 0)
    for _, char := range log {
        if char == oldRune {
            foo = append(foo, newRune)
        } else {
            foo = append(foo, char)
        }
    }
    return string(foo)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
    if utf8.RuneCountInString(log)  <= limit {
        return true
    } 
    return false
}
