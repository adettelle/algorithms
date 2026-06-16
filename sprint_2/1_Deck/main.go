// https://contest.yandex.ru/contest/22781/run-report/163000734/
/*
-- ПРИНЦИП РАБОТЫ --
Реализована структура данных Дек, максимальный размер которого определяется
заданным числом. Элемент можно добавить в начало или конец дека.
Используется кольцевой буфер.

Если в деке уже находится максимальное число элементов, элемент не будет добавлен.
Удалить можно первый или последний элемент. При этом он выводится на печать.
Если дек был пуст, то ничего не удаляется.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Доказательство корректности дека (двусторонней очереди) на кольцевом буфере
базируется на строгом математическом инварианте структуры данных и методе полной индукции.
Оно гарантирует, что все операции вставки/удаления выполняются за время O(1)
без нарушения порядка элементов и переполнения.

Пусть физический массив имеет размер N.
Индексы варьируются от 0 до N - 1.
Указатели:
H (Head) — индекс первого элемента,
T (Tail) — индекс первого свободного слота.
Дек является "закольцованным", поэтому все операции с индексами вычисляются
по модулю размера массива: index_next = (index+-1) mod N

Корректность алгоритма обеспечивается следующими правилами, которые остаются истинными
до и после выполнения любой операции:
Проверка на пустоту (isEmpty): Дек пуст, если указатели размер стека равен нулю.
Проверка на заполненность (isFull): Дек заполнен, размер и емкость стека равны.
Количество элементов: увеличивается / уменьшается при добавлении / удалении элемента.

Докажем корректность операций:
- Добавление push_front:
H = (H - 1 + N) mod N, затем запись в data[H].
Если дек был не полон, head смещается назад в пустую ячейку, формула размера
корректно увеличивается на 1.
- Добавление вправо push_back:
Запись в data[T], затем T = (T + 1) mod N.
Запись происходит в зарезервированный слот. Указатель Tail указывает на следующий
свободный слот.
- Удаление (pop_front / pop_back):
Перед извлечением проверяется условие пустоты. При удалении мы сдвигаем
соответствующий указатель. Так как размер вычисляется динамически,
удаленные элементы просто перезаписываются новыми, предотвращая выход за границы.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Все операции (push_front, push_back, pop_front, pop_back) выполняются за O(1),
потому что доступ к элементам слайса по индексу происходит за О(1), вычисление новых позиций
head и tail тоже за О(1).
Обработка всех команд происходит за O(n), где n - количество команд.

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Дек хранится в массиве фиксированного размера (память O(capacity)), также используются несколько полей
(size, capacity, head, tail), которые требуют памяти O(1).
Память для хранения списка команд - O(n).
Итог: O(capacity + n)

*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	n := scanInt(scanner)
	cap := scanInt(scanner)

	commands := make([]string, n)
	for i := range n {
		commands[i] = scanLine(scanner)
	}

	deck := NewDeck(cap)
	deck.ExecuteCmds(commands)
}

func (d *Deck) ExecuteCmds(commands []string) {
	for _, command := range commands {
		res, err := d.ExecuteCmd(command)
		if err != nil {
			fmt.Println("error")
		} else {
			if res != nil {
				fmt.Println(*res)
			}
		}
	}
}

func (d *Deck) ExecuteCmd(command string) (*int, error) {
	cmd := strings.Split(command, " ")
	val := 0
	var res *int
	var err error

	if len(cmd) > 1 {
		val, err = strconv.Atoi(cmd[1])
		if err != nil {
			return res, err
		}
	}

	var x int
	switch cmd[0] {
	case "pop_back":
		x, err = d.PopBack()
		res = &x
	case "pop_front":
		x, err = d.PopFront()
		res = &x
	case "push_front":
		err = d.PushFront(val)
	case "push_back":
		err = d.PushBack(val)
	}
	return res, err
}

func scanInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	strN := scanner.Text()
	x, _ := strconv.Atoi(strN)
	return x
}

func scanLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

type Deck struct {
	data     []int
	head     int
	tail     int
	capacity int
	size     int
}

func NewDeck(cap int) *Deck {
	return &Deck{
		data:     make([]int, cap),
		head:     0,
		tail:     0,
		capacity: cap,
		size:     0,
	}
}

func (d *Deck) IsEmpty() bool {
	return d.size == 0
}

func (d *Deck) isFull() bool {
	return d.size == d.capacity
}

// Хвост всегда указывает на первую свободную для записи ячейку
// голова — на элемент, добавленный в очередь раньше всех остальных
func (d *Deck) PushFront(value int) error {
	if d.isFull() {
		return fmt.Errorf("deck is empty")
	}
	newHead := (d.head - 1 + d.capacity) % d.capacity
	d.data[newHead] = value
	d.head = newHead

	d.size++
	return nil
}

func (d *Deck) PushBack(value int) error {
	if d.isFull() {
		return fmt.Errorf("deck is full")
	}

	d.data[d.tail] = value
	d.tail = (d.tail + 1) % d.capacity
	d.size++

	return nil
}

func (d *Deck) PopFront() (int, error) {
	if d.IsEmpty() {
		return 0, fmt.Errorf("error")
	}
	val := d.data[d.head]
	d.data[d.head] = 0
	d.head = (d.head + 1) % d.capacity
	d.size--
	// чтобы при size=0 head и tail вставали на нулевую позицию - это необязательно
	// if d.size == 0 {
	// 	d.head = 0
	// 	d.tail = 0
	// }

	return val, nil
}

func (d *Deck) PopBack() (int, error) {
	if d.IsEmpty() {
		return 0, fmt.Errorf("error")
	}

	idxToDelete := (d.capacity + d.tail - 1) % d.capacity
	val := d.data[idxToDelete]
	d.data[idxToDelete] = 0
	d.tail = (idxToDelete) % d.capacity
	d.size--
	// чтобы при size=0 head и tail вставали на нулевую позицию - это необязательно
	// if d.size == 0 {
	// 	d.head = 0
	// 	d.tail = 0
	// }
	return val, nil
}
