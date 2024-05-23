/*
Read 2 integer values and store them in variables A and B. Make the sum of A and B,
assigning your result to variable X. Print X as per the example shown below.
Do not present any message other than what is being specified and do not forget to print the end of line after the result,
otherwise you will receive "Presentation Error".

Input
The input contains 2 integer values.

Exit
Print the message "X = " (capital letter X) followed by the value of the variable X and the end of the line.
Make sure there is a space before and after the equal sign, as shown in the example below.
*/
package main

import "fmt"

func main() {
	//DECLARAÇÃO DAS VARIAVEIS
	var x, y, total int32

	//ENTRADA DE DADOS
	fmt.Scan(&x)
	fmt.Scan(&y)

	//CALCULO DO TOTAL
	total = x + y

	//SAÍDA
	fmt.Printf("X = %v", total)
}
