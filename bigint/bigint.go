package bigint

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

/*
An implementation of BigInt using Go.
This structure stores large integers as strings essentially an array of digits 0 to 9 representing the larger value.
Includes methods for:
	* comparing, adding, and subtracting two BigInt values.
	* dividing a BigInt value by an integer value
	* multiplying two BigInt values together
	* negating a BigInt
	* creating a string with the sign of a number prepended
	* divide two BigInt values to get the quotient and remainder
*/

// BigInt type holds the value as a string and a boolean to track if the integer is negative (true) or positive (false)
type BigInt struct {
	value    string
	negative bool
}

// Create a BigInt from a string of numbers
// returns the BigInt or an error when the input string is invalid
func NewBigInt(v string) (*BigInt, error) {

	// check for a minus sign
	if v[0] == '-' {

		// remove minus sign
		newV := v[1:]

		// check string is all digits
		if !checkDigits(newV) {
			return &BigInt{}, errors.New("not a valid big int string")
		}

		b := &BigInt{value: newV, negative: true}
		return b, nil
	}

	// check string is all digits
	if !checkDigits(v) {
		return &BigInt{}, errors.New("not a valid big int string")
	}

	b := &BigInt{value: v, negative: false}
	return b, nil
}

// ToString returns a string of the BigInt value
// appends a minus sign for negative values
func (b *BigInt) ToString() string {
	if b.negative {
		return fmt.Sprintf("-%s", b.value)
	}
	return b.value
}

// Negate changes the sign of a BigInt
func (b *BigInt) Negate() {
	b.negative = !b.negative
}

// checkDigits returns true if all values in a string are between 0 and 9
func checkDigits(s string) bool {
	for _, char := range s {
		if char < '0' || char > '9' {
			// return false at the first non digit
			return false
		}
	}
	return true
}

// getValue returns the string of the BigInt
func (b *BigInt) getValue() string {
	return b.value
}

// setValue sets the value of the BigInt
func (b *BigInt) setValue(v string) {
	b.value = v
}

// runes returns the value of the BigInt as a slice of runes
func (b *BigInt) runes() []rune {
	return []rune(b.value)
}

// length returns the length of the BigInt string
func (b *BigInt) length() int {
	return len(b.value)
}

// compareValues checks if one string value is larger than another
// compareValues does not account for positive or negative values
// -1 if b < x
// 0 if b == x
// 1 if b > x
func (b *BigInt) compareValues(x *BigInt) int {
	if b.length() > x.length() {
		return 1
	}
	if b.length() < x.length() {
		return -1
	}

	thisRunes := b.runes()
	xRunes := x.runes()

	for i := 0; i < b.length(); i++ {
		bInt := int(thisRunes[i] - '0')
		xInt := int(xRunes[i] - '0')

		if bInt > xInt {
			return 1
		}
		if bInt < xInt {
			return -1
		}
	}
	return 0
}

// CompareTo checks if BigInt b is larger, smaller, or equal to BigInt x
// CompareTo accounts for positive and negative values
// -1 if b < x
// 0 if b == x
// 1 if b > x
func (b *BigInt) CompareTo(x *BigInt) int {

	// b is positive x is positive
	if !b.negative && !x.negative {

		// The longer value is larger (more positive)
		if b.length() > x.length() {
			return 1
		}
		if b.length() < x.length() {
			return -1
		}

		thisRunes := b.runes()
		xRunes := x.runes()

		// compare digits of strings. Larger value is larger (more positive)
		for i := 0; i < b.length(); i++ {
			bInt := int(thisRunes[i] - '0')
			xInt := int(xRunes[i] - '0')
			if bInt > xInt {
				return 1
			}
			if bInt < xInt {
				return -1
			}
		}
		return 0
	}

	// b is positive x is negative
	if !b.negative && x.negative {
		return 1
	}

	// b is negative x is negative
	if b.negative && x.negative {
		// The longer value is smaller (more negative)
		if b.length() > x.length() {
			return -1
		}
		if b.length() < x.length() {
			return 1
		}

		thisRunes := b.runes()
		xRunes := x.runes()

		// compare digits of strings. Larger value is smaller (more negative)
		for i := 0; i < b.length(); i++ {
			bInt := int(thisRunes[i] - '0')
			xInt := int(xRunes[i] - '0')
			if bInt > xInt {
				return -1
			}
			if bInt < xInt {
				return 1
			}
		}
		return 0
	}

	// b is negative x is positive
	if b.negative && !x.negative {
		return -1
	}

	return 0
}

// equalLengths adds zeros to make each BigInt string the same length
func equalLengths(x *BigInt, y *BigInt) (string, string) {
	var xVal, yVal string
	// Same length. Don't need to add zeros
	if x.length() == y.length() {
		yVal = y.getValue()
		xVal = x.getValue()
		// Pad y value with zeroes
	} else if x.length() > y.length() {
		xVal = x.getValue()
		yVal = strings.Repeat("0", (x.length()-y.length())) + y.getValue()
		// Pad x value with zeroes
	} else {
		xVal = strings.Repeat("0", (y.length()-x.length())) + x.getValue()
		yVal = y.getValue()
	}

	return xVal, yVal
}

// reverse inverts the order of runes in a string
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// adder is a helper function for adding two BigInt values together
func (b *BigInt) adder(addend *BigInt) {
	// ensure both values are equal length
	x, y := equalLengths(addend, b)

	// track addition carrying
	overflow := 0

	// solution is the resulting sum
	var solution strings.Builder

	// iterate through the value of each big int from right to left.
	for i := len(x) - 1; i >= 0; i-- {
		xInt := int(x[i] - '0')
		yInt := int(y[i] - '0')
		// Add the two digits of each BigInt and any overflow from the previous addition
		digitSum := xInt + yInt + overflow
		// using mod 10 ensure the value is from 0 to 9
		// append the added digit to the sum
		fmt.Fprintf(&solution, "%d", (digitSum % 10))
		// calculate any future overflow either a 0 or 1 if there is carrying
		overflow = digitSum / 10
	}

	// Once all digits are iterated check for any carrying.
	s := "1"
	if overflow == 0 {
		s = ""
	}
	// Reverse the string to account for append order and add any last carrying
	b.value = s + reverse(solution.String())

}

// Add one BigInt value with another
func (b *BigInt) Add(x *BigInt) {

	// b is smaller
	if b.compareValues(x) == -1 {
		// b is positive x is negative
		if !b.negative && x.negative {
			xValue := x.value
			x.subtractor(b)
			b.value = x.value
			b.negative = true
			x.value = xValue
			return
		}

		// b is positive x is positive
		if !b.negative && !x.negative {
			b.adder(x)
			b.negative = false
			return
		}

		// b is negative x is negative
		if b.negative && x.negative {
			b.adder(x)
			b.negative = true
			return
		}

		// b is negative x is positive
		if b.negative && !x.negative {
			xValue := x.value
			x.subtractor(b)
			b.value = x.value
			b.negative = true
			x.value = xValue
			return
		}
	}

	// b is larger
	if b.compareValues(x) == 1 {
		// b is positive x is negative
		if !b.negative && x.negative {
			b.subtractor(x)
			b.negative = false
			return
		}

		// b is positive x is positive
		if !b.negative && !x.negative {
			b.adder(x)
			b.negative = false
			return
		}

		// b is negative x is negative
		if b.negative && x.negative {
			b.adder(x)
			b.negative = true
			return
		}

		// b is negative x is positive
		if b.negative && !x.negative {
			b.subtractor(x)
			b.negative = true
			return
		}
	}

	// b and x are equal
	if b.compareValues(x) == 0 {
		// b is positive x is positive
		if !b.negative && !x.negative {
			b.adder(x)
			b.negative = false
			return
		}

		// b is positive x is negative
		if !b.negative && x.negative {
			b.value = "0"
			b.negative = false
			return
		}

		// b is negative x is positive
		if b.negative && !x.negative {
			b.value = "0"
			b.negative = false
			return
		}

		// b is negative x is negative
		if b.negative && x.negative {
			b.adder(x)
			b.negative = true
		}
	}
}

// Subtract one BigInt from another
func (b *BigInt) Subtract(x *BigInt) {
	// b is smaller
	if b.compareValues(x) == -1 {
		// b is positive x is negative
		if !b.negative && x.negative {
			b.adder(x)
			b.negative = false
			return
		}

		// b is positive x is positive
		if !b.negative && !x.negative {
			xValue := x.value
			x.subtractor(b)
			b.value = x.value
			x.value = xValue
			b.negative = true
			return
		}

		// b is negative x is negative
		if b.negative && x.negative {
			xValue := x.value
			x.subtractor(b)
			b.value = x.value
			x.value = xValue
			b.negative = false
			return
		}

		// b is negative x is positive
		if b.negative && !x.negative {
			b.adder(x)
			b.negative = true
			return
		}
	}

	// b is larger
	if b.compareValues(x) == 1 {
		// b is positive x is negative
		if !b.negative && x.negative {
			b.adder(x)
			b.negative = false
			return
		}

		// b is positive x is positive
		if !b.negative && !x.negative {
			b.subtractor(x)
			b.negative = false
			return
		}

		// b is negative x is negative
		if b.negative && x.negative {
			b.subtractor(x)
			b.negative = true
			return
		}

		// b is negative x is positive
		if b.negative && !x.negative {
			b.adder(x)
			b.negative = true
			return
		}
	}

	// b and x are equal
	if b.compareValues(x) == 0 {
		// b is positive x is positive
		if !b.negative && !x.negative {
			b.value = "0"
			b.negative = false
			return
		}

		// b is positive x is negative
		if !b.negative && x.negative {
			b.adder(x)
			b.negative = false
			return
		}

		// b is negative x is positive
		if b.negative && !x.negative {
			b.adder(x)
			b.negative = true
			return
		}

		// b is negative x is negative
		if b.negative && x.negative {
			b.value = "0"
			b.negative = false
			return
		}
	}
}

// subtractor is a helper method for subtracting two big int
func (b *BigInt) subtractor(subtrahend *BigInt) {
	// Pad subtrahend with zeros to make each BigInt equal length
	x, y := equalLengths(subtrahend, b)
	overflow := 0

	// sb holds the final value
	var sb strings.Builder

	// work from right to left along the subtrahend (x)
	for i := len(x) - 1; i >= 0; i-- {
		// Get values of subtrahend and minuend
		xInt := int(x[i] - '0')
		// Add overflow from previous difference to account for any borrowing
		yInt := int(y[i]-'0') + overflow

		// overflow can be 0 or -1
		// it is -1 when a 1 is needed from the next 10s place
		overflow = 0
		// if the subtrahend is larger than minuend then a borrow is needed
		if yInt < xInt {
			overflow = -1
			yInt += 10
		}
		// append the difference to the solution builder
		fmt.Fprintf(&sb, "%d", yInt-xInt)
	}

	// Remove any trailing zeroes, before reversing, they come as a result of borrowing
	solution := sb.String()
	for solution[len(solution)-1] == '0' {
		solution = solution[0 : len(solution)-1]
	}

	result := reverse(solution)
	b.value = result
}

// multiply one BigInt with another
func (b *BigInt) Multiply(x *BigInt) {
	xRunes := x.runes()
	number := b.getValue()
	b.value = "0"

	// multiply each digit of x (the multiplier) BigInt with all digits of b BigInt (the multiplicand)
	for i := 0; i < x.length(); i++ {
		// digit from i starting at the right (powerOf10 is zero at the right)
		multiplier := int(xRunes[len(xRunes)-1-i] - '0')
		// sb is used to build the product through recursive steps
		var sb strings.Builder
		// Use the multiplyByIntHelper to recursively
		newB, _ := NewBigInt(multiplyByIntHelper(multiplier, number, i, 0, &sb))
		// Add the previous product to b
		b.adder(newB)
	}

	// Find the resulting sign for b
	// both values are negative or positive results in a positve
	if b.negative == x.negative {
		b.negative = false
	} else {
		// opposite signs of b and x results in a negative
		b.negative = true
	}

}

// multiplyByInt is a helper method for multiplying two BigInt
func multiplyByIntHelper(x int, number string, powerOf10 int, overflow int, sb *strings.Builder) string {
	// Base recursive step
	// When the last digit of multiplicand is reached pad with zeroes to match the current powerOf10
	if len(number) == 0 {
		s := strconv.Itoa(overflow)
		if overflow == 0 {
			s = ""
		}
		return s + reverse(sb.String()) + strings.Repeat("0", powerOf10)
	}
	// get the multiplicant from the remaining number
	multiplicand := int(number[len(number)-1] - '0')
	// multiply the values by integer and add any overflow from previous multiplications
	product := multiplicand*x + overflow
	// append the value to the builder as a digit from 0 to 9 using mod
	fmt.Fprintf(sb, "%d", (product % 10))
	// calculate any new overflow using integer division
	newOverflow := product / 10
	// recursively call the next mutiplicand by reducing number by 1
	return multiplyByIntHelper(x, number[0:len(number)-1], powerOf10, newOverflow, sb)
}

// divideByInt will divide a BigInt by an integer value
// will lose precision because of integer division
func (b *BigInt) DivideByInt(divisor int) {
	// cannot divide by zero
	if divisor == 0 {
		log.Fatal("cannot divide by zero")
	} else {
		// sb is the quotient
		var sb strings.Builder
		dividend := b.getValue()
		overflow := 0
		// work from left to right along the dividend
		for i := 0; i < len(dividend); i++ {
			// digit is the number to be divided
			// add any overflow from the previous division
			digit := overflow*10 + int(dividend[i]-'0')
			// add the divided digit to the quotient
			fmt.Fprintf(&sb, "%d", digit/divisor)
			// overflow is the remainder from the division and used in the next division
			overflow = digit % divisor
		}

		// Remove any leading zeroes resulting from carrying
		solution := sb.String()
		for solution[0] == '0' {
			solution = solution[1:]
		}

		b.value = solution
	}

}

// Divide a BigInt value by another b / x
// Integer division tells us how many times x fits into b
// returns a string of any remainder.
func (b *BigInt) Divide(x *BigInt) string {

	remainder := "0"

	if b.compareValues(x) == 1 {
		// b is larger
		// subtractions tracks the number of times x fits into b
		subtractions := 0
		// While the numerator is larger or equal subtract the denominator
		for b.compareValues(x) >= 0 {
			b.subtractor(x)
			subtractions++
		}
		// remainder is the remainder value of b after subtractions
		remainder = b.value

		// The result of integer division is the number of times x fits into b
		b.value = strconv.Itoa(subtractions)

	} else if b.compareValues(x) == -1 {
		// b is smaller	EX 10 / 100 = 0, Remainder 10
		remainder = b.value
		b.value = "0"

	} else if b.compareValues(x) == 0 {
		// b and x are equal 10 / 10 = 1, Remainder 0
		b.value = "1"
	}

	// Find the resulting sign for b
	// both values are negative or positive results in a positve
	if b.negative == x.negative {
		b.negative = false
	} else {
		// opposite signs of b and x results in a negative
		b.negative = true
	}

	return remainder
}

func main() {

	num1, _ := NewBigInt("0")
	num2, _ := NewBigInt("-932423400")

	num3, _ := NewBigInt("10003")
	num4, _ := NewBigInt("-1000")

	num3.Add(num4)
	fmt.Println(num3.ToString())

	num1.Add(num2)
	fmt.Println(num1.ToString())

	num5, _ := NewBigInt("-900")
	num6, _ := NewBigInt("-1000")

	num5.Subtract(num6)
	fmt.Println(num5.ToString())
	fmt.Println(num6.ToString())

	num7, _ := NewBigInt("500")
	fmt.Println(num7.ToString())
	num7.Negate()
	fmt.Println(num7.ToString())

	fmt.Println(num5.CompareTo(num3))

	num8, _ := NewBigInt("1000")
	num8.DivideByInt(10)
	fmt.Println(num8.ToString())

	num9, _ := NewBigInt("725")
	num10, _ := NewBigInt("-725")
	result := num9.CompareTo(num10)
	fmt.Println(result)

	num11, _ := NewBigInt("-10")
	num12, _ := NewBigInt("-10")
	num11.Multiply(num12)
	fmt.Println(num11.ToString())

	num13, _ := NewBigInt("1000000")
	num14, _ := NewBigInt("1120")
	remainder := num13.Divide(num14)
	fmt.Println("Divide:", num13.ToString(), "Remainder", remainder)
}
