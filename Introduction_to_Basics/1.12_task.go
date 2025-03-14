/*
По данному трехзначному числу определите, все ли его цифры различны.

Формат входных данных
На вход подается одно натуральное трехзначное число.

Формат выходных данных
Выведите "YES", если все цифры числа различны, в противном случае - "NO".

Sample Input 1:

237
Sample Output 1:

YES
Sample Input 2:

117
Sample Output 2:

NO
*/
package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	a := num / 100
	b := (num / 10) % 10
	c := num % 10
	if a != b && a != c && b != c {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
