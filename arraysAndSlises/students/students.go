package students

import "fmt"

func Average(s []float64) {
	var sum float64

	for _, v := range s {
		sum += v
	}

	fmt.Println(sum / float64(len(s)))
}
