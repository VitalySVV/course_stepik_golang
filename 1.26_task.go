/*
На ввод подаются пять целых чисел, которые записываются в массив. Однако эта часть программы уже написана.

Вам нужно написать фрагмент кода, с помощью которого можно найти и вывести максимальное число в этом массиве.

Sample Input:
2
3
56
45
21

Sample Output:
56
*/
package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	maxNumber := array[0]
	for _, element := range array {
		if element > maxNumber {
			maxNumber = element
		}
	}
	fmt.Println(maxNumber)
}
