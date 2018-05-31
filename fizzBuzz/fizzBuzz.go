package fizzBuzz

import (
	"strconv"
)

func say(number int) string {

	var fizzBuzz = map[int]string{
		3: "Fizz",
		5: "Buzz",
	}
	str := ""
	array := []int{3, 5}
	for i := 0; i < len(array); i++ {
		if number%array[i] == 0 {
			str += fizzBuzz[array[i]]
		}

	}
	if str != "" {
		return str
	}
	return strconv.Itoa(number)

}
