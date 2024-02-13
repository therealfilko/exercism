package raindrops

import (
	"strconv"
)

func Convert(number int) string {
    foo := map[int]string{ 3: "Pling", 5: "Plang", 7: "Plong"}
    result := ""
    for teiler, klang := range foo {
        if number % teiler == 0 {
            result += klang
        } 
    }
    if len(result) == 0 {
        return strconv.Itoa(number)
    } 
    return result
}
