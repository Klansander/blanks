package main

import "fmt"

func main() {
	minMaxDifference(90)
}

func minMaxDifference(num int) int {
	mass := make([]int, 0, 5)
	massmin := make([]int, 0, 5)
	massmax := make([]int, 0, 5)
	k := 9
	for num != 0 {
		mass = append(mass, num%10)
		massmin = append(massmin, num%10)
		massmax = append(massmax, num%10)
		if num%10 != 9 {
			k = num % 10
		}
		num /= 10

	}
	for i := len(massmax) - 1; i >= 0; i-- {
		if k == massmin[i] {
			massmin[i] = 0
			massmax[i] = 9
		}

	}
	res1 := 0
	res2 := 0
	fmt.Println(massmin, massmax)
	for i := len(massmax) - 1; i >= 0; i-- {

		res1 += massmin[i]
		res1 *= 10
		res2 += massmax[i]
		res2 *= 10

	}
	res1 /= 10
	res2 /= 10
	fmt.Println(res1, res2)
	return res2 - res1
}
