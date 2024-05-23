/*
In this exercise we use the functions Math.Cos(x) and Math.Sin(x) which calculate,
respectively, the cosine and sine of the angle presented to it within the parameter.
After the necessary imports, the variable ‘number’ was declared and
data entry by the user, who types the Angle that wants to discover the sine and cosine.
Then the cosine and sine of the desired angle are calculated and printed,
using the functions math.Cos(angle) and math.Sin(angle) as print parameters
from fmt.PrintIn().
*/
package main

import (
	"fmt"
	"math"
)

func main() {

	var angulo float64

	fmt.Print("Digite o ângulo desejado:\n")
	fmt.Scan(&angulo)

	cos := math.Cos(angulo)
	sin := math.Sin(angulo)

	fmt.Println("Cosseno de", angulo, "é:", cos)

	fmt.Println("Seno de", angulo, "é:", sin)

}
