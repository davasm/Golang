/*
Beecrowd 1002
The formula for calculating the area of a circle is: area = π. radius2. Considering for this problem that π = 3.14159:

- Calculate the area, squaring the radius value and multiplying by π.

Input
The input contains a floating point value (double precision), in this case, the variable radius.

Exit
Display the message "A=" followed by the value of the area variable, as per the example below,
with 4 places after the decimal point. Use double precision variables.
Like all problems, don't forget to print the end of line after the result,
otherwise you will get "Presentation Error".
*/
package main

import (
	"fmt"
	"math"
)

func main() {

	const PI float64 = 3.14159
	var raio float64

	fmt.Scan(&raio)

	area := PI * math.Pow(raio, 2)

	fmt.Printf("A=%.4f\n", area)
}
