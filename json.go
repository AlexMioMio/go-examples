package main

import "fmt"
import "encoding/json"

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	slcE := []string{"abc", "def", "xyz"}
	slcD, _ := json.Marshal(slcE)
	fmt.Println(string(slcD))

	mapA := map[string]int{"abc": 1, "def": 2}
	mapC, _ := json.Marshal(mapA)
	fmt.Println(string(mapC))

	res1 := &Response1{
		Page:   1,
		Fruits: []string{"abc", "def"}}
	resJ, _ := json.Marshal(res1)
	fmt.Println(string(resJ))

	res2 := &Response2{
		Page:   2,
		Fruits: []string{"def", "ghi"}}
	resJ2, _ := json.Marshal(res2)
	fmt.Println(string(resJ2))

	a := 5
	switch {
	case a > 1:
	case a > 2:
	case a > 3:
		fmt.Println("shit2")
	default:
		fmt.Println("fuck3")
	}

}