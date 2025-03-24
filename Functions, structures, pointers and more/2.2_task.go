/*Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

Входные данные

Вводится четыре числа.

Выходные данные

Необходимо вернуть из функции наименьшее из 4-х данных чисел

Sample Input:

4 5 6 7
Sample Output:

4*/
func minimumFromFour() int {
	var k, x int
	k = 0
	var min int
	fmt.Scan(&min)
	for fmt.Scan(&x); k < 3; fmt.Scan(&x) {
		if x < min {
			min = x
		}
		k++
	}
	return min
}