package main

import (
	"encoding/json"
	"fmt"
	"os"
	_ "os"
)

const (
	a = iota + 1
	_
	b
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func testFunc() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	arr := []string{"s", "sd", "ds"}
	val := arr[3]
	fmt.Println(val)
}

func b1() {
	fmt.Println(1211)
}
func a1() {

	defer b1()
	c1()
}

func c1() {
	arr := []string{"s", "sd", "ds"}
	val := arr[3]
	fmt.Println(val)
}

func change(a *int) {
	fmt.Println(&a)
	t := *a * 2
	a = &t
	fmt.Println(&a)
}

type model struct {
	Name string `json:"test"`
	Age  int    `json:"age"`
}

func main() {

	//err := Foo()
	//fmt.Println(err, err == nil)
	//
	//testFunc()
	//
	//a := make([]int, 2)
	//a = append(a, 2)
	//a[0] = 1
	//a = append(a, 4)
	//fmt.Println(a)
	//
	////a1()
	//
	//s := "sadaффы@"
	//fmt.Println(len(s), len([]byte(s)), len([]rune(s)))
	//bb := 2
	//change(&bb)
	//fmt.Println(bb)

	//task := make(chan int, 1)
	//task <- 1
	//fmt.Println(122)
	//
	//fmt.Println(t)

	datr := `{ "name":"sdsa","age":23}`
	var model model
	err := json.Unmarshal([]byte(datr), &model)
	if err != nil {
		fmt.Println("wsqew")
	}
	fmt.Println(model)
}
