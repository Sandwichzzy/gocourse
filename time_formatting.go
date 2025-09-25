package main

import (
	"fmt"
	"time"
)


func main() {

	// Mon Jan 2 15:04:05 MST2006
	layout := "2006-01-02T15:04:05Z07:00"
	str := "2024-07-04T14:30:18Z"

	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	fmt.Println(t) //2024-07-04 14:30:18 +0000 UTC

	str1 := "Jul 03, 2024 03:18 PM"
	layout1 := "Jan 02, 2006 03:04 PM"

	t1, _ := time.Parse(layout1, str1)
	fmt.Println(t1) //2024-07-03 15:18:00 +0000 UTC
}
