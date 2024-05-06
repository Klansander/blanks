package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	_ "github.com/xuri/excelize/v2"
	"strconv"
)

func main() {
	f, err := excelize.OpenFile("1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	err = f.SetSheetName("Sheet1", "a")
	if err != nil {
		fmt.Println(err)
	}
	a := "CDEFGHIJKL"
	aa := "OPQRSTUVWX"
	c := 0
	arr := make([]int, 0, 600)
	for i := 3; i <= 62; i++ {
		for _, j := range a {
			cell, err := f.GetCellValue("a", string(j)+strconv.Itoa(i))
			if err != nil {
				fmt.Println(err)
				return
			}
			//fmt.Println(cell)
			b, _ := strconv.Atoi(cell)
			arr = append(arr, b)
			c++
		}
	}

	c = 0
	for i := 3; i <= 62; i++ {
		for _, l := range aa {
			err := f.SetCellValue("a", string(l)+strconv.Itoa(i), arr[c])
			if err != nil {
				fmt.Println(err, l)
				return
			}

			c++
		}
	}
	// Получить значение из ячейки по заданному имени и оси листа

	for j, _ := range arr {
		arr[j] -= 2
	}

	map1 := make(map[int]int)
	map2 := make(map[int]int)
	for i := len(arr) - 1; i >= 0; i-- {
		map1[arr[i]] = i
	}
	for i, v := range map1 {
		map2[v] = i
	}

	for j, _ := range arr {
		arr[j] = map1[arr[j]]
	}

	for j, _ := range arr {

		if arr[j]%2 == 0 {
			arr[j] -= 4
		} else {
			arr[j] += 4
		}

	}

	v := []int{1, 2, 3, 4, 5}
	c = 0
	for c < 600 {
		for _, j := range v {

			arr[c] += j
			c++

		}
	}
	fmt.Println(arr)

	c = 599
	for i := 3; i <= 62; i++ {
		for _, l := range aa {
			err := f.SetCellValue("a", string(l)+strconv.Itoa(i), arr[c])
			if err != nil {
				//	fmt.Println(err, l)
				return
			}

			c--
		}
	}

	c = 0
	for i := 3; i <= 62; i++ {
		for _, l := range aa {
			err := f.SetCellValue("a", string(l)+strconv.Itoa(i), arr[c])
			if err != nil {
				//fmt.Println(err, l)
				return
			}

			c++
		}
	}

	main()

	c = 0
	for i := 3; i <= 62; i++ {
		for _, l := range aa {
			err := f.SetCellValue("a", string(l)+strconv.Itoa(i), arr[c])
			if err != nil {
				//fmt.Println(err, l)
				return
			}

			c++
		}
	}
	err = f.Save()
	if err != nil {

		return
	}
}
