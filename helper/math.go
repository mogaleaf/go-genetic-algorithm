package helper

import "math"

func StandardDeviationInt(numbers []int, total int) float64 {
	mean := total / len(numbers)
	count := 0.0
	for _, iter := range numbers {
		diffsq := math.Pow(float64(iter-mean), 2)
		count += diffsq
	}
	meansq := count / float64(len(numbers))
	return math.Sqrt(meansq)
}
