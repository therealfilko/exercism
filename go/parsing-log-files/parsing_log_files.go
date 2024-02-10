package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
    foo := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    return foo.MatchString(text)
}

func SplitLogLine(text string) []string {
    return regexp.MustCompile(`<[~*=-]*>`).Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
    pattern := `"(?i).*password.*"`
    re := regexp.MustCompile(pattern)

    count := 0
    for _, line := range lines {
        if re.FindString(line) != "" {
            count++
        }
    }
    return count
}

func RemoveEndOfLineText(text string) string {
    re := regexp.MustCompile(`end-of-line\d+`)
    cleanedText := re.ReplaceAllString(text, "")
    return cleanedText
}

func TagWithUserName(lines []string) []string {
    re := regexp.MustCompile(`User\s+(\S+)`)
    var result []string

    for _, line := range lines {
        match := re.FindStringSubmatch(line)
        if match != nil {
            // Füge [USR] und den Benutzernamen am Anfang der Zeile ein
            taggedLine := "[USR] " + match[1] + " " + line
            result = append(result, taggedLine)
        } else {
            // Füge die unveränderte Zeile hinzu, wenn kein Benutzername gefunden wurde
            result = append(result, line)
        }
    }

    return result
}

