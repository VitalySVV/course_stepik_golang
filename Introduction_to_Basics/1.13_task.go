/*
Дано неотрицательное целое число. Найдите и выведите первую цифру числа.

Формат входных данных
На вход дается натуральное число, не превосходящее 10000.

Формат выходных данных
Выведите одно целое число - первую цифру заданного числа.

Sample Input:

1234
Sample Output:

1
*/
package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	switch {
	case num >= 10000:
		fmt.Println(num / 10000)
	case num >= 1000:
		fmt.Println(num / 1000)
	case num >= 100:
		fmt.Println(num / 100)
	case num >= 10:
		fmt.Println(num / 10)
	default:
		fmt.Println(num)
	}
}
