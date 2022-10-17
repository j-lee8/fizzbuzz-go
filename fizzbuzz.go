package main

import (
	"fmt"
	"strconv"
)

func main() {
	fizzBuzz(50);
}

func fizzBuzz(n int) []string {
	var fizzBuzzArr []string;

	for i := 0; i < n; i++ {

		divisibleBy3 := (i % 3 == 0);
        divisibleBy5 := (i % 5 == 0);

		if (divisibleBy3 && divisibleBy5) {
			fizzBuzzArr = append(fizzBuzzArr, "FizzBuzz");
		} else if (divisibleBy3) {
			fizzBuzzArr = append(fizzBuzzArr, "Fizz");
		} else if (divisibleBy5) {
			fizzBuzzArr = append(fizzBuzzArr, "Buzz");
		} else {
			fizzBuzzArr = append(fizzBuzzArr, strconv.Itoa(i));
		}
	}

	fmt.Println("Result: ", fizzBuzzArr);

	return fizzBuzzArr;
}