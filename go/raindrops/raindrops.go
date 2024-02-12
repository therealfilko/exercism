package raindrops

import "fmt"

func Convert(number int) string {
    foo := map[int]string{ 3: "Pling", 5: "Plang", 7: "Plong"}
    for i := 0; i <= 100; i++ {
        for teiler, klang := range foo {
            fmt.Println("Zahl: %d", i)
            if i % teiler == 0 {
                fmt.Println("Zahl: %s", klang)
            }
        }
    }
}
