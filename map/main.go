package main

import "fmt"

func main()  {
	m1 := make(map[string]string)

	m1["test"] = "hola"
	fmt.Println(m1["test"])

	mutateMap(m1)
	fmt.Println(m1["test"])
}

func mutateMap(m map[string]string) {
	m["test"] = "hello"
}