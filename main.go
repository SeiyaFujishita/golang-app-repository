package main

import "fmt"

func main() {
	m := make([]map[string]int, 0)
	newMap1 := make(map[string]int)
	newMap2 := make(map[string]int)
	newMap3 := make(map[string]int)
	newMap1["key"] = 1
	newMap2["key"] = 2
	newMap3["key"] = 3
	m = append(m, newMap1)
	m = append(m, newMap2)
	m = append(m, newMap3)

	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}
}
