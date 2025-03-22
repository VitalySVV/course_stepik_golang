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
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	for index, element := range arr {
		if index%2 == 0 {
			fmt.Print(element, " ")
		}
	}
}
