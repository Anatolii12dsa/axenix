package main

import "fmt"

func average(numbers []float64) float64 {
    total := 0.0
    for _, num := range numbers {
        total += num
    }
    return total / float64(len(numbers))
}

func main() {
    numbers := []float64{3.5, 6.2, 8.9, 10.1, 4.8}
    avg := average(numbers)
    fmt.Println("Среднее значение:", avg)
}