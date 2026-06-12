package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	pattern, _ := reader.ReadString('\n')
	sequence, _ := reader.ReadString('\n')

	if isSubsequence(sequence, pattern) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

func isSubsequence(sequence string, pattern string) bool {
	j := 0

	for i := 0; i < len(sequence); i++ {
		if sequence[i] == pattern[j] {
			j++
		}
		if j == len(pattern) {
			return true
		}
	}
	return false
}

func isSubsequenceRecursion(sequence string, pattern string) bool {
	// база, условие выхода:
	// 1. когда просмотрены все эл-ты pattern, когда j == len(pattern)
	// 2. Если строка sequence закончилась раньше, чем строка pattern, Возвращаем True.
	// значит найти оставшиеся символы уже невозможно. Возвращаем False.

	// Шаг рекурсии: Функция вызывает сама себя с новыми данными
	// 1. Сравниваем первые ещё не рассмотренные символы строк pattern и sequence.
	// 2. Если символы совпадают, значит этот символ подпоследовательности найден.
	// Переходим к следующему символу в обеих строках.
	// 3. Если символы не совпадают, то текущий символ строки sequence нам не подходит.
	// Пропускаем его и переходим к следующему символу только в строке sequence.
	// 4. Рекурсивно повторяем процесс.
	return check(pattern, sequence, 0, 0)
}

func check(pattern, sequence string, j, i int) bool {
	// Все символы pattern найдены
	if j == 0 {
		return true
	}

	// sequence закончилась раньше
	if i == len(sequence) {
		return false
	}

	// Символы совпали — двигаемся по обеим строкам
	if pattern[j] == sequence[i] {
		return check(pattern, sequence, j+1, i+1)
	}
	// Символы не совпали — пропускаем символ в t
	return check(pattern, sequence, j, i+1)
}
