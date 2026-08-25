package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, capacity int
	fmt.Fscan(reader, &n, &capacity)

	Backpacks := make([]Backpack, n)

	for i := 0; i < n; i++ {
		var w, p int
		fmt.Fscan(reader, &w, &p)
		Backpacks[i] = Backpack{Weight: w, Price: p}
	}

	items := FindTotalWeight(Backpacks, capacity)
	fmt.Println(len(items))
	for i := 0; i < len(items); i++ {
		fmt.Print(items[i], " ")
	}
}

type Backpack struct {
	Weight int
	Price  int
}

// A(s, n) есть максимальная стоимости предметов, которые можно уложить в рюкзак
// максимальной вместимости n, если можно использовать только первые
// s предметов из заданных k.
func FindTotalWeight(backpacks []Backpack, capacity int) []int {
	n := len(backpacks)
	// r - количество доступных предметов
	//  с - текущая вместимость рюкзака
	// dp[r][c] - максимальная стоимость, которую можно получить,
	// используя первые r предметов, если вместимость рюкзака равна c.
	dp := make([][]int, n+1)

	// Нулевая строка очень важна: если предметов нет, максимальная стоимость всегда 0.
	for r := 0; r < n+1; r++ {
		dp[r] = make([]int, capacity+1)
	}

	// r - номер строки, то есть сколько предметов мы сейчас рассматриваем
	// r = 1 -> рассматриваем backpacks[0] и т.д.
	for r := 1; r <= n; r++ {
		item := backpacks[r-1]

		// c - это исключительно текущая вместимость рюкзака
		for c := 0; c <= capacity; c++ {
			// пока не берём текущий предмет
			dp[r][c] = dp[r-1][c]

			if c >= item.Weight { // хватает ли места для текущего предмета?
				// Если взять текущий предмет, сколько максимально можно получить?
				value := dp[r-1][c-item.Weight] + item.Price

				// Выбираем лучший вариант
				// Вариант 1 - не брать предмет dp[r][c]
				// Вариант 2 - взять предмет: dp[r-1][c-item.Weight] + item.Price
				if value > dp[r][c] {
					dp[r][c] = value
				}
			}
		}
	}

	for r := 0; r < n+1; r++ {
		fmt.Println(dp[r])
	}
	// Восстанавливаем выбранные предметы
	items := make([]int, 0)
	c := capacity

	for r := n; r >= 1; r-- {
		if dp[r][c] != dp[r-1][c] {
			items = append(items, r) // нумерация предметов с 1
			c -= backpacks[r-1].Weight
		}
	}
	// dp[k][w] - максимальная стоимость
	return items
}
