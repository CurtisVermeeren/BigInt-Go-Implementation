# Go BigInt

A BigInt implementation written in Go as practice. 
BigInt values are stored as strings in this implementation. 
A boolean flag called negative is used to indicate if a value is positive or negative.

To create a new BigInt use `NewBigInt(v string)` this method creates a new BigInt object from a string of numbers. It will check that the v string is a valid input.
Create values with a string like `"1234054"` to mark a new value as negative prepend a dash `"-5009054"`, This will set `BigInt.negative` to true. 

### Methods called on BigInt b

These methods are called on a BigInt object called b in this example. The value of b is altered depending on the operation. 

`ToString()` Return a string of the BigInt with a minus sign appended for negative values.

`Negate()` Reverses the polarity of the BigInt. Positive values will be negated. Negative values will become positive.

`CompareTo(x *BigInt)` is used to compate two BigInt values b and x. Returns -1 if b < x. Returns 0 of b == x. Returns 1 if b > x.

`Add(x *BigInt)` Add the value of x to BigInt b.

`Subtract(x *BigInt)` Subtract the value of x from BigInt b.

`Multiply(x *BigInt)` Multiply BigInt b by the value of x.

`DivideByInt(divisor int)` Divide BigInt b by the value of the integer divisor.

`Divide(x *BigInt)` Divide BigInt b by the value of x.

