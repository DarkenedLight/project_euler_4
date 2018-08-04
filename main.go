/*
Problem:
A palindromic number reads the same both ways. The largest palindrome made from
the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.

Notes:
this could have been done more easily by using the math package or much more
easily by using strings or slices, but i wanted to challenge myself by only
working with ints and the fmt package
*/

package main

import "fmt"

//receives an int arguement and returns a bool saying if it is a palindrome
func isPalin(num int) bool {
	//initialize variables
	var first int
	var last int
	var skip int
	for {
		//num is a single digit, which is itself a palindrome
		if (num/10) == 0 && skip < 0 {
			break
		}

		//special case where after removing the first digit, the new first
		//digit is 0. ex, 10101 % 10000 returns 101, but logically you are
		//working with 0101
		if skip > 0 {
			first = 0
		} else {
			first = getFirstDigit(num)
		}
		last = num % 10

		//first and last digit dont match, not a palindrome
		if first != last {
			return false
		}

		//if the special case occured, there is not need to remove the first
		//digit, decrement number of times to skip calling "removeFitstDigit"
		if skip > 0 {
			skip--
		} else {
			length := getLen(num)
			num = removeFirstDigit(num)
			lengthAfter := getLen(num)

			if (length - 1) != lengthAfter {
				skip = (length - lengthAfter) - 1
			}
		}
		//remove last digit
		num /= 10
	}
	return true
}

//continuously removes the last digit until one digit is left, being the first,
//and then returns that digit
func getFirstDigit(num int) int {
	var first int
	for {
		//one digit left, first digit found
		if (num / 10) == 0 {
			break
		}
		//remove last digit
		first = num / 10
		num /= 10
	}
	return first
}

//pretty much performs num % 10^(len(num)-1), or re-written this performs
//num % int(10^int(log10(num)))
func removeFirstDigit(num int) int {
	length := getLen(num)
	dom := 1
	for i := 1; i < length; i++ {
		dom *= 10
	}
	return num % dom
}

//returns the length of an int by continuously removing the last digit
//and keeping count as you go
func getLen(num int) int {
	var count int
	for {
		//one digit left, icnrement count and break
		if (num / 10) == 0 {
			count++
			break
		}
		//increment and remove last digit
		count++
		num /= 10
	}
	return count
}

func main() {
	var maxX int
	var maxY int
	for x := 999; x >= 0; x-- {
		for y := 999; y >= 0; y-- {
			if isPalin(x * y) {
				if (x * y) > (maxX * maxY) {
					maxX = x
					maxY = y
				}
			}
		}
	}

	fmt.Println(maxX*maxY, "=", maxX, "x", maxY)
}
