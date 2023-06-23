package main

import (
	"fmt"
	"math"

	"github.com/CurtisVermeeren/bigint-go-implementation/bigint"
)

func main() {

	// Creating 4 different BigInt
	num1, _ := bigint.NewBigInt("0")
	num2, _ := bigint.NewBigInt("-932423400")

	num3, _ := bigint.NewBigInt("10003")
	num4, _ := bigint.NewBigInt("-1000")

	// Creating Invalid BigInt
	_, err := bigint.NewBigInt("hello")
	if err != nil {
		fmt.Println("Creating BigInt with \"hello\" : " + err.Error() + "\n")
	}

	_, err = bigint.NewBigInt("90no0")
	if err != nil {
		fmt.Println("Creating BigInt with \"90no0\" : " + err.Error() + "\n")
	}

	_, err = bigint.NewBigInt("+123")
	if err != nil {
		fmt.Println("Creating BigInt with \"+1234\" : " + err.Error() + "\n")
	}

	// Addition examples
	fmt.Println("Adding: " + num4.ToString() + " to " + num3.ToString())
	num3.Add(num4)
	fmt.Println("Result: " + num3.ToString() + "\n")

	fmt.Println("Adding: " + num2.ToString() + " to " + num1.ToString())
	num1.Add(num2)
	fmt.Println("Result: " + num1.ToString() + "\n")

	num1, _ = bigint.NewBigInt("3456")
	num2, _ = bigint.NewBigInt("56789")
	fmt.Println("Adding: " + num1.ToString() + " to " + num2.ToString())
	num1.Add(num2)
	fmt.Println("Result: " + num1.ToString() + "\n")

	// Create more BigInt
	num5, _ := bigint.NewBigInt("-900")
	num6, _ := bigint.NewBigInt("-1000")

	// Subtraction examples
	fmt.Println("Subtracting: " + num6.ToString() + " from " + num5.ToString())
	num5.Subtract(num6)
	fmt.Println("num5: " + num5.ToString() + "\n")

	num1, _ = bigint.NewBigInt("345234526")
	num2, _ = bigint.NewBigInt("-56789")
	fmt.Println("Subtracting: " + num2.ToString() + " from " + num1.ToString())
	num1.Subtract(num2)
	fmt.Println("num1: " + num1.ToString() + "\n")

	num1, _ = bigint.NewBigInt("345234526")
	num2, _ = bigint.NewBigInt("333567893")
	fmt.Println("Subtracting: " + num2.ToString() + " from " + num1.ToString())
	num1.Subtract(num2)
	fmt.Println("num1: " + num1.ToString() + "\n")

	// Negating num7
	num7, _ := bigint.NewBigInt("500")
	fmt.Println("num7: " + num7.ToString())
	num7.Negate()
	fmt.Println("num7 after negation: " + num7.ToString() + "\n")

	// Comparting two values
	fmt.Println("Comparing: " + num5.ToString() + " with " + num3.ToString())
	printComparison(num5, num3)

	num9, _ := bigint.NewBigInt("725")
	num10, _ := bigint.NewBigInt("-725")
	fmt.Println("Comparing: " + num9.ToString() + " with " + num10.ToString())
	printComparison(num9, num10)

	// Dividing using Integer Division
	num8, _ := bigint.NewBigInt("1002")
	fmt.Println("num8: " + num8.ToString())
	num8.DivideByInt(10)
	fmt.Println("num8 divided by 10 using Integer Division: " + num8.ToString() + "\n")

	// Multiplying two BigInt values
	num11, _ := bigint.NewBigInt("-10")
	num12, _ := bigint.NewBigInt("-10")
	fmt.Println("Multiplying: " + num11.ToString() + " with " + num12.ToString())
	num11.Multiply(num12)
	fmt.Println("Result : " + num11.ToString() + "\n")

	// Dividing two BigInt values
	num13, _ := bigint.NewBigInt("1000000")
	num14, _ := bigint.NewBigInt("1120")
	fmt.Println("Dividing: " + num13.ToString() + " by " + num14.ToString())
	remainder := num13.Divide(num14)
	fmt.Println("Result: " + num13.ToString() + " Remainder: " + remainder + "\n")

	fmt.Printf("The int min and max values in Go from math.MinInt and math.MaxInt are: %d and %d\n", math.MinInt, math.MaxInt)
	num1, _ = bigint.NewBigInt("9223372036854775808")
	num2, _ = bigint.NewBigInt("9223372036854775808")
	fmt.Println("Adding: " + num1.ToString() + " to " + num2.ToString())
	num1.Add(num2)
	fmt.Println("Result: " + num1.ToString() + "\n")

}

func printComparison(x, y *bigint.BigInt) {
	if x.CompareTo(y) == 1 {
		fmt.Println(x.ToString(), "is larger than", y.ToString()+"\n")
	} else if x.CompareTo(y) == -1 {
		fmt.Println(x.ToString(), "is smaller than", y.ToString()+"\n")
	} else {
		fmt.Println(x.ToString(), "is equal to", y.ToString()+"\n")
	}
}
