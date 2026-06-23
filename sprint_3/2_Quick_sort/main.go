// https://contest.yandex.ru/contest/23815/run-report/163238308/

/*
-- ПРИНЦИП РАБОТЫ --
Для хранения данных о студентах выбран тип данных - структура.
Реализована функция - компаратор biggerComparator для сравнения двух студентов
по баллам, штрафам и именам.
Для сортировки студентов используется алгоритм быстрой сортировки, но поскольку
по условию реализация сортировки не может потреблять O(n) дополнительной памяти
для промежуточных данных, применяется модификация быстрой сортировки,
называемой "in-place".

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Рассмотрим вызов функции quickSortWithComparator при n > 1.

Выбираем опорный элемент (pivot), в нашем случае последний элемент массива.
Во время прохода по массиву каждый элемент сравнивается с pivot при помощи компаратора.
Если элемент должен находиться раньше pivot, он переносится в левую часть массива;
иначе остаётся в правой части.
После завершения цикла опорный элемент помещается между двумя частями массива.
Следовательно, каждый элемент слева от pivot меньше него в соответствии
с условиями сравнения. Таким образом, после разбиения опорный элемент
оказывается на своей окончательной позиции.
Далее алгоритм рекурсивно сортирует левую и правую части.

Получается, что весь массив оказывается отсортированным в требуемом порядке.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
На каждом уровне рекурсии выполняется разбиение массива за O(n).
В среднем случае опорный элемент делит массив примерно пополам,
поэтому глубина рекурсии составляет O(log n).
Тогда общая сложность равна O(nlogn).
В худшем случае (например, если массив уже упорядочен относительно
выбранного опорного элемента) разбиение получается неравномерным:
размеры подмассивов равны 0 и n - 1. Тогда глубина рекурсии составляет O(n),
а суммарная сложность O(n^2).

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Алгоритм выполняет сортироку на месте, не создавая дополнительных массивов,
кроме фиксированного числа переменных.
Пространственная сложность алгоритма без учёта рекурсивного стека O(1).
Но дополнительная память расходуется на стек рекурсивных вызовов.
В среднем случае глубина рекурсии составляет O(log n), а в худшем случае — O(n).
*/
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	n := readInt(scanner)
	students := readArray(scanner, n)

	quickSortWithComparator(students, Less)
	for _, student := range students {
		fmt.Println(student.name)
	}
}

func Less(a, b Student) bool {
	if a.points != b.points { // >
		return a.points > b.points
	}

	if a.fine != b.fine {
		return a.fine < b.fine
	}

	return a.name < b.name
}

func quickSortWithComparator(arr []Student, less func(Student, Student) bool) {
	if len(arr) < 2 {
		return
	}

	pivot := arr[rand.Intn(len(arr)-1)]

	left := 0
	right := len(arr) - 1

	for left <= right {
		for left <= right && less(arr[left], pivot) {
			left++
		}
		for left <= right && less(pivot, arr[right]) {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	if right > 0 {
		quickSortWithComparator(arr[:right+1], less)
	}
	if left < len(arr) {
		quickSortWithComparator(arr[left:], less)
	}
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	res, _ := strconv.Atoi(scanner.Text())
	return res
}

func readArray(scanner *bufio.Scanner, n int) []Student {
	students := make([]Student, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		listStr := strings.Split(scanner.Text(), " ")
		student := getStudent(listStr)
		students[i] = student
	}

	return students
}

type Student struct {
	name   string
	points int
	fine   int
}

func getStudent(listStr []string) Student {
	points, _ := strconv.Atoi(listStr[1])
	fine, _ := strconv.Atoi(listStr[2])

	student := Student{
		name:   listStr[0],
		points: points,
		fine:   fine,
	}
	return student
}
