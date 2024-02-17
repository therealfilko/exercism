package hamming

import "errors"

func Distance(a, b string) (int, error) {
    // Wir sollen checken nach sequenz, welche unterschiede sind und diese dann zählen und zurückgeben
    count := 0
    if len(a) != len(b) {
        return -1, errors.New("DNA Stränge sind nicht gleich lang")
    }
    for i := range a {
        if a[i] != b[i] {
            count += 1
        }
    }
    return count, nil
}
