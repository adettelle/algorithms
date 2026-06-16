// https://contest.yandex.ru/contest/22781/run-report/163001392/
/*
-- ПРИНЦИП РАБОТЫ --
Для вычисления значения выражения, записанного в обратной польской нотации,
используются две структуры: словарь для хранения операторов и их функций
и Стэк для хранения операндов.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
После обработки первых k токенов выражения стек содержит
значения всех полностью вычисленных подвыражений, соответствующих этим токенам.
При этом верхние элементы стека являются результатами последних вычисленных подвыражений
и могут использоваться для дальнейших операций.

Рассмотрим обработку очередного токена.
Если токен является числом, алгоритм помещает его в стек.
Число само по себе является корректным подвыражением.
Если токен является оператором, то по свойству корректной записи в обратной польской нотации
два верхних элемента стека являются результатами двух последних вычисленных подвыражений.
Алгоритм извлекает их из стека, применяет оператор и помещает результат обратно в стек.
Полученное значение является результатом нового подвыражения.

После обработки всех токенов все возможные операции выполнены,
а стек содержит результаты вычисленных подвыражений.
Согласно условию задачи, если в стеке осталось несколько чисел,
необходимо вывести только верхний элемент.
Алгоритм извлекает верхний элемент стека и выводит его.
Этот элемент является результатом последнего вычисленного подвыражения,
то есть требуемым ответом.

Следовательно, алгоритм корректно вычисляет значение выражения в обратной польской нотации.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Пусть n — количество токенов во входном выражении.
Алгоритм последовательно обрабатывает каждый токен ровно один раз.
Для каждого токена выполняется константное число операций:
- проверка наличия оператора в словаре — O(1);
- добавление элемента в стек — O(1);
- извлечение элемента из стека — O(1);
- выполнение арифметической операции — O(1).
Итог: суммарная временная сложность алгоритма составляет O(n).

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
В худшем случае все токены являются операндами, поэтому в стеке может
одновременно находиться до n чисел.
Дополнительная память, используемая словарём операторов, постоянна
и не зависит от размера входных данных.
Итог: пространственная сложность алгоритма составляет O(n).
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	elems := strings.Split(line, " ")

	operators := map[string]func(a, b int) int{
		"+": sum,
		"-": subtraction,
		"/": division,
		"*": multiplication,
	}

	result, err := reversedPolishNotation(elems, operators)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func reversedPolishNotation(elems []string,
	operators map[string]func(a, b int) int) (int, error) {
	stack := NewStack()

	for _, elem := range elems {
		if f, ok := operators[elem]; ok {
			// operator
			arg2, err := stack.pop()
			if err != nil {
				return 0, err
			}
			arg1, err := stack.pop()
			if err != nil {
				return 0, err
			}
			res := f(arg1, arg2)
			stack.push(res)
		} else {
			// operand
			num, err := strconv.Atoi(elem)
			if err != nil {
				return 0, fmt.Errorf("unable convert to integer: %w", err)
			}
			stack.push(num)
		}
	}
	result, err := stack.pop()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func subtraction(a, b int) int {
	return a - b
}

func sum(a, b int) int {
	return a + b
}

func multiplication(a, b int) int {
	return a * b
}

// математическое целочисленное деление (округление всегда происходит вниз)
func division(a, b int) int {
	div := float64(a) / float64(b)

	return int(math.Floor(div))
}

type Stack struct {
	stack []int
}

func NewStack() *Stack {
	return &Stack{stack: []int{}}
}

func (s *Stack) push(val int) {
	s.stack = append(s.stack, val)
}

func (s *Stack) pop() (int, error) {
	size := len(s.stack)
	if size == 0 {
		return 0, fmt.Errorf("empty stack")
	}
	res := s.stack[size-1]
	s.stack = s.stack[:size-1]
	return res, nil
}
