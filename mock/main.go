package main

import (
	"fmt"
	"math"
)

func main() {
	//a1 := make([]int, 0)
	//a1 = append(a1, []int{1, 2, 3, 4, 5}...)
	//fmt.Println(cap(a1))
	//a2 := append(a1, 6, 5, 44, 55)
	//a3 := append(a1, 7)
	//a2[1] = -1
	//a3[2] = -2
	//fmt.Println(a1, a2, a3)

	//a1 := make([]int, 0)
	//a1 = append(a1, []int{1, 2, 3, 4, 5}...)
	//a2 := append(a1, 6)
	//a3 := append(a1, 7)
	//
	//fmt.Println(a1, a2, a3)

	//first := []int{10, 20, 30, 40}
	//second := make([]*int, len(first))
	//for i, v := range first {
	//	second[i] = &v
	//	fmt.Println(&i, &v)
	//}
	//fmt.Println(*second[0], *second[1])
	fmt.Println(int('d') - int('t'))
	s := "eduktdb"
	k := 14
	j := 0
	s1 := ""
	for i := 1; i < len(s); i++ {

		if len(s1) == 0 {
			if math.Abs(float64(int(s[i])-int(s[i-1]))) <= float64(k) {
				s1 += string(s[i-1])
				s1 += string(s[i])
				j += 2
			}
		} else if int(math.Abs(float64(int(s[i])-int(s1[j-1])))) <= k {
			fmt.Println(int(s[i])-int(s1[j-1]), string((s[i])))
			s1 += string(s[i])
			j++
		} else if len(s1) != 0 && int(math.Abs(float64(int(s[i])-int(s1[j-1])))) > k {
			fmt.Println(int(s[i])-int(s1[j-1]), string((s[i])))
			s1 += string(s[i])
			j++
		}
		fmt.Println(string(s1))

	}
}

//func handle() error {
//	return &myError{text: "error"}
//}
//
//type myError struct {
//	text string
//}
//
//func (m myError) Error() string {
//	return m.text
//}
