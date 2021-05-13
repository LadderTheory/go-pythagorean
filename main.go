package main

import (
	"fmt"
	"math"
)

// formulas
// a = h^2 - k^2;
// b = 2hk;
// c = h^2 + k^2;
// sum formula: a + b + c = 2(h^2 + hk)
// formula to find k from h & s(sum): -h + (s/2h)
// abc = (h^2 - k^2)(2hk)(h^2 + k^2) = 2h^5k - 2hk^5

func main() {
	start := 1
	end := 10000

	sum_map := []map[int][][]int{}

	for i := start; i <= end; i += 1 {
		hk := hk_from_sum(i)
		if len(hk) > 2 {
			sum_map = append(
				sum_map,
				map[int][][]int{
					i: hk,
				},
			)
		}
	}

	for _, v := range sum_map {
		for k, e := range v {
			fmt.Println(k, len(e))
		}
	}
}

func hk_from_sum(sum1 int) [][]int {
	var hk = [][]int{}
	for h := 2; true; h += 1 {
		k := find_k(h, sum1)
		if k >= 1 {
			sum2 := sum(h, k)
			if k < h && sum2 == sum1 {
				//fmt.Println(h, k)
				hk = append(
					hk,
					[]int{h, k},
				)
			}
		} else {
			break
		}
	}

	return hk
}

func find_k(h int, s int) int {
	// -h + (s/2h) = k
	return (0 - h) + (s / (2 * h))
}

func sum(h int, k int) int {
	// -h + (s/2h)
	return 2 * ((h * h) + (h * k))
}

func product(h int, k int) int {
	// 2h^5k - 2hk^5
	return (2 * int(math.Pow(float64(h), 5)) * k) - (2 * h * int(math.Pow(float64(k), 5)))
}
