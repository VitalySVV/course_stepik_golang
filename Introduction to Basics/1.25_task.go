/*
Напишите программу, принимающая на вход число
N(N≥4), а затем N целых чисел-элементов среза. На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.

Sample Input:
5
41 -231 24 49 6

Sample Output:
49
*/
package main

import "fmt"

func main() {
	var n int
	var arr []int
	if fmt.Scan(&n); n >= 4 {
		arr = make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}
	}
	fmt.Println(arr[3])
}
