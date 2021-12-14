package main

import (
	"fmt"
	"strings"
)

type BigInt struct {
  value string
}

// Create a BigInt from a string of numbers
func newBigInt(v string) *BigInt {
  b := &BigInt{value: v}
  return b
}

// getValue returns the string of the BigInt
func (b *BigInt) getValue() string{
  return b.value;
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

// compareTo checks if BigInt b is larger, smaller, or equal to BigInt x
// -1 if b < x
// 0 if b == x
// 1 if b > x
func (b *BigInt) compareTo(x *BigInt) int{
  if (b.length() > x.length()) {
    return 1
  }
  if (b.length() < x.length()) {
    return -1
  }

  thisChars := b.runes()
  xChars := x.runes()

  for i := 0; i < b.length(); i++ {
    xInt := int(thisChars[i] - '0')
    bInt := int(xChars[i] - '0')
    if (bInt > xInt) {
      return 1
    }
    if (bInt < xInt) {
      return -1
    }
  }
  return 0
}

// equalLengths adds zeros to make each BigInt string the same length
func equalLengths(x *BigInt, y *BigInt) (string, string) {
  var xVal, yVal string
  // Same length. Don't need to add zeros
  if (x.length() == y.length()) {
    yVal = y.getValue()
    xVal = x.getValue()
  // Pad y value with zeroes
  } else if (x.length() > y.length()) {
    xVal = x.getValue()
    yVal = strings.Repeat("0", (x.length() - y.length())) + y.getValue()
  // Pad x value with zeroes
  } else {
    xVal = strings.Repeat("0", (y.length() - x.length())) + x.getValue()
    yVal = y.getValue()
  }

  return xVal, yVal
}

// Reverse inverts the order of runes in a string
func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

// plus adds one BigInt value with another
func (b *BigInt) plus(xVal *BigInt) {
  // ensure both values are equal length
  x, y := equalLengths(xVal, b)
  
  // track addition carrying
  overflow := 0

  // solution is the resulting sum
  var solution strings.Builder

  // iterate through the value of each big int from right to left.
  for i := len(x) - 1; i >= 0; i-- {
    xInt := int(x[i] - '0')
    yInt := int(y[i] - '0')
    // Add the two digits of each BigInt and any overflow from the previous addition
    digitSum := xInt + yInt + overflow;
    // addInt is the value appended to the sum
    // using mod 10 ensure the value is from 0 to 9
    addInt := digitSum % 10
    // append the added digit to the sum
    fmt.Fprintf(&solution, "%d", addInt)
    // calculate any future overflow either a 0 or 1 if there is carrying
    overflow = digitSum / 10;
  }

  // Once all digits are iterated check for any carrying.
  s := "1"
  if overflow == 0 {
    s = ""
  }
  // Reverse the string to account for append order and add any last carrying
  b.value = s + Reverse(solution.String())
}



func main() {

	num := newBigInt("666666666666666666634555555555555555555555555555555555555555533466")
  num2 := newBigInt("33333333555555555555555555555555555555555555543333")

  num.plus(num2)
  fmt.Println(num.value)

  fmt.Println(9 / 10)
 
}
