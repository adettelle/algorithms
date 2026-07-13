// https://contest.yandex.ru/contest/24414/run-report/163586273/

/*
-- ПРИНЦИП РАБОТЫ --
Создается хеш-таблица - структура HashTable (список списков элементов Pair,
которые хранят пару: ключ, значение).
Количество бакетов задаётся исходя из заданного n,
поскольку все команды могут, например, складывать значения в таблицу.
Далее реализованы методы над этой структурой: put, get и delete
в соответсвии с заданием.
Чтобы не превысить количество бакетов, индекс бакета в методах задается
как модуль от остатка от деления ключа на количество бакетов.
Если ключ уже имеется, то операция put заменит соответствующее ему значение.
Операция get выводит найденное значение или "None" при отсутсвии
значения по ключу. Операция delete удаляет найденное значение
или выводит "None" при отсутсвии значения по ключу.


-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Рассмотрим выполняемые операции.

Операция put(key, value).
Вычисляется номер корзины (bucket) по хеш-функции от ключа.
Затем последовательно просматриваются все элементы этой корзины.
Если элемент с данным ключом уже существует, его значение заменяется на новое.
Если такого ключа нет, в конец корзины добавляется новая пара (key, value).
Таким образом, после выполнения операции в таблице существует
ровно одна запись с данным ключом, содержащая переданное значение.

Операция get(key).
По ключу вычисляется номер корзины, в которой может находиться искомый элемент.
Затем выполняется последовательный просмотр элементов этой корзины.
Если найден элемент с данным ключом, возвращается связанное с ним значение.
Если элемент не найден, возвращается "None".

Операция delete(key).
По ключу определяется корзина, после чего выполняется поиск элемента.
Если элемент найден, запоминается его значение, после чего
он удаляется из корзины заменой на последний элемент
и уменьшением длины среза. Остальные элементы таблицы
продолжают храниться. Если элемент отсутствует, возвращается "None".
Следовательно, после удаления ключ больше не содержится в таблице,
а возвращаемое значение соответствует удалённому элементу.

Каждая операция работает только с корзиной, определяемой хеш-функцией,
и обрабатывает все элементы этой корзины.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Каждая операция выполняет вычисление хеша за O(1) и просматривает
только одну корзину.
В худшем случае, если все ключи попали в одну корзину,
каждая операция требует просмотра всех элементов и работает за O(n),
где n - количество хранимых элементов

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Хеш-таблица хранит все n элементов словаря.
Также создаётся массив корзин (hashTable), размером (buckets = n / capacity),
пусть будет m.
Пространственная сложность равна O(n+m)
Поскольку число корзин выбирается пропорционально числу элементов,
(m = n / capacity), то итоговая пространственная сложность равна O(n).
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	hash := NewHashTable(n)

	for i := 0; i < n; i++ {
		cmd := readCommand(reader)
		switch cmd.Cmd {
		case "put":
			hash.put(cmd.Key, cmd.Value)
		case "get":
			res, ok := hash.get(cmd.Key)
			if !ok {
				fmt.Fprintln(writer, "None")
			} else {
				fmt.Fprintln(writer, res)
			}
		case "delete":
			res, ok := hash.delete(cmd.Key)
			if !ok {
				fmt.Fprintln(writer, "None")

			} else {
				fmt.Fprintln(writer, res)
			}
		}
	}
}

type Pair struct {
	Key   int
	Value int
}

type HashTable struct {
	buckets [][]Pair
}

func NewHashTable(n int) HashTable {
	capacity := 10
	buckets := 32

	if n > 32 {
		buckets = n / capacity
	}
	hashTable := make([][]Pair, buckets)
	hash := HashTable{buckets: hashTable}

	return hash
}

func (h *HashTable) findPairIndex(bucket []Pair, key int) (int, bool) {
	for i, x := range bucket {
		if x.Key == key {
			return i, true
		}
	}

	return 0, false
}

func (h *HashTable) put(key, value int) {
	bucketIdx := Abs(key % len(h.buckets))

	pairIndex, ok := h.findPairIndex(h.buckets[bucketIdx], key)
	if ok {
		h.buckets[bucketIdx][pairIndex].Value = value
	} else {
		h.buckets[bucketIdx] = append(h.buckets[bucketIdx], Pair{
			Key:   key,
			Value: value,
		})
	}
}

func (h *HashTable) get(key int) (int, bool) {
	bucketIdx := Abs(key % len(h.buckets))

	var value int

	pairIndex, ok := h.findPairIndex(h.buckets[bucketIdx], key)
	if ok {
		value = h.buckets[bucketIdx][pairIndex].Value
		return value, ok
	}

	return 0, ok
}

func (h *HashTable) delete(key int) (int, bool) {
	bucketIdx := Abs(key % len(h.buckets))

	var value int

	pairIndex, ok := h.findPairIndex(h.buckets[bucketIdx], key)
	if ok {
		value = h.buckets[bucketIdx][pairIndex].Value
		last := len(h.buckets[bucketIdx]) - 1
		h.buckets[bucketIdx][pairIndex] = h.buckets[bucketIdx][last]
		h.buckets[bucketIdx] = h.buckets[bucketIdx][:last]
		return value, ok
	}

	return 0, ok
}

type Command struct {
	Cmd   string
	Key   int
	Value int
}

func readCommand(reader *bufio.Reader) Command {
	var cmd string
	var key, value int
	fmt.Fscan(reader, &cmd, &key)

	if cmd == "put" {
		fmt.Fscan(reader, &value)
	}

	command := Command{Cmd: cmd, Key: key, Value: value}

	return command
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
