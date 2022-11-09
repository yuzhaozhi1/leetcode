package main

import "fmt"

/*
	count := 0
	for i := low; i <= high; i++ {
		fmt.Println(i)
		if i%2 != 0 {
			count++
		}
	}
	return count
*/

func countOdds(low int, high int) int {
	// count := 0
	// start := low
	// end := high
	// for start < end {
	// 	if start%2 != 0 {
	// 		count++
	// 	}
	// 	if end%2 != 0 {
	// 		count++
	// 	}
	// 	start++
	// 	end--
	// }
	// return count

	count := 0
	for i := low; i <= high; i++ {
		if i%2 != 0 {
			count++
		}
	}
	return count
}
func main() {
	odds := countOdds(3, 7)
	fmt.Println(odds)
}
