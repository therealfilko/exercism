package raindrops

import (
	"strconv"
)

func Convert(number int) string {
    foo := map[int]string{ 3: "Pling", 5: "Plang", 7: "Plong"}
    bar := []int{3, 5, 7}
    result := ""
    for _, teiler := range bar {
        if number % teiler == 0 {
            result += foo[teiler]
        } 
    }
    if len(result) == 0 {
        return strconv.Itoa(number)
    } 
    return result
}
