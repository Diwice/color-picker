package main

import (
	"fmt"
	"math"
)

func main() {
	num_1, num_2, num_3 := float64(200)/255.0, 0.055, 1.055
	fmt.Println(num_1 + num_2)
	fmt.Println((num_1 + num_2)/num_3)
	fmt.Println(math.Pow((num_1 + num_2)/num_3, 2.4))
}
