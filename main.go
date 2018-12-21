package main

import (
		"fmt"
		)

func main() {
	var n int
	fmt.Scan(&n)

	var s string
	fmt.Scan(&s)

	maru := "O"
	batsu := "X"
	count := 0
	i := 0
	for i + 1 < n {
		if s[i:i+2] == maru + batsu || s[i:i+2] == batsu + maru {
			count++
			i += 2
		}else{
			i++
		}
	}
	fmt.Println(count)
}
