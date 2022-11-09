package main

import "fmt"

func main() {
	fmt.Println(isValid(""))
}

func isValid(s string) bool {
	data1 := make(map[uint8]int, 1000)
	data2 := make(map[uint8]int, 1000)
	start := 0
	end := len(s) - 1
	mid := len(s) / 2
	for start < end {
		i := s[start]
		j := s[end]
		if _, ok := data[i]; ok {
			data[i]++
		} else {
			data[i] = 1
		}
		if _, ok := data[j]; ok {
			data[j]++
		} else {
			data[j] = 1
		}
		start++
		end--
	}

	if data['('] != data[')'] {
		return false
	}
	if data['{'] != data['}'] {
		return false
	}
	if data['['] != data[']'] {
		return false
	}
	return true

}
