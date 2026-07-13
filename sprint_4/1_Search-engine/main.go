// https://contest.yandex.ru/contest/24414/run-report/163583530/

/*
-- ПРИНЦИП РАБОТЫ --
По заданию на вход получаем n документов, каждый из которых
представляет собой текст из слов. Индексируем все слова во всех документах.

В итоге имеем map, где ключом является слово, а в значении находися
список сктруктур, где хранятся позиция документа и количество слова в нём.

Пример документов:
cat dog (1)
cat     (2)
cat cat (3)

Итоговая карта релевантности:
cat: [{1; 1} {2; 1} {3; 2}
dog: [{1; 1}]

Далее получаем релевантность докуметов.
Релевантность документа оценивается следующим образом:
для каждого уникального слова из запроса берётся число его вхождений в документ,
полученные числа для всех слов из запроса суммируются.
Итоговая сумма и является релевантностью документа.
Чем больше сумма, тем больше документ подходит под запрос.

Таким образом, для запроса "cat cat dog" повторяющиеся слова будут удалены,
запрос будет "cat dog".

После вычисления релевантности всех подходящих документов формируется
список из не более чем пяти наиболее релевантных документов.
Для этого документы по одному добавляются в список top, который всегда
поддерживается отсортированным по убыванию релевантности,
а при равной релевантности — по возрастанию номера документа.
Если после вставки количество элементов превышает пять,
последний элемент удаляется.

В результате для каждого запроса выводятся документы в требуемом порядке.


-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
При построении индекса (makeMap) для каждого слова и каждого документа
сохраняется количество вхождений этого слова в документ.

Перед обработкой запроса из него удаляются повторяющиеся слова,
поэтому каждое уникальное слово запроса учитывается ровно один раз,
как требуется по условию.

Для каждого слова запроса алгоритм (getRelevance) просматривает все документы,
содержащие это слово, и прибавляет к релевантности документа
число вхождений данного слова. Таким образом, после обработки
всех слов запроса для каждого документа вычисляется сумма количеств вхождений
всех уникальных слов запроса.

После обработки каждого документа в функции collectTop
список top содержит не более пяти лучших документов среди уже обработанных и
отсортирован по требуемому порядку. Новый документ вставляется
перед первым документом, который уступает ему по релевантности или,
при равной релевантности, имеет больший номер.
Если после вставки размер списка превышает пять, последний элемент удаляется.
После обработки всех документов в top остаются ровно пять
наиболее релевантных документов (или меньше, если подходящих документов
меньше пяти).

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
N — число документов
K - максимальная длина документа в символах
M — количество запросов
L - максимальная длина запроса в символах

Построение индекса makeMap: проходимся по каждому документу: O(N),
обработка одного документа O(K), получается O(K*N)

При обработке запроса сначала сортируются слова запроса (slices.Sort)
для последующего удаления повторов (slices.Compact), что требует O(L log L).
Перебираем все слова в запросе O(L).
После этого для каждого уникального слова запроса просматривается
список документов, содержащих данное слово.
Каждое слово может встретиться во всех N документах,
поэтому вычисление релевантностей в худшем случае требует O(L*N) операций.

Поддержание списка из пяти наиболее релевантных документов выполняется
за константное время для каждого обработанного документа,
поскольку его размер никогда не превышает пяти.
Поэтому эта операция не изменяет итоговую оценку.

Таким образом, обработка одного запроса имеет временную сложность O(L*logL + L*N),
а обработка всех запросов: O(M*(L*logL + L*N)).

Итоговая временная сложность всей программы: O(K*N + M*(L*logL + L*N)).


-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Индекс хранит для каждого слова список документов, в которых оно встречается,
вместе с количеством вхождений.
Память, необходимая для хранения индекса, составляет O(N*K).

При обработке одного запроса создаётся отображение relevance, в котором
для каждого документа, содержащего хотя бы одно слово запроса,
накапливается его релевантность. В худшем случае в него могут попасть
все документы, поэтому он требует O(N) памяти.

Список из пяти лучших документов имеет постоянный размер, следовательно,
требует O(1) памяти.

Итоговая пространственная сложность O(N + N*K)
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	docs := make([][]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		docs[i] = strings.Fields(scanner.Text())
	}

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	const queriesToShow = 5

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	index := makeMap(docs)

	for i := 0; i < m; i++ {
		scanner.Scan()
		query := strings.Fields(scanner.Text())

		slices.Sort(query)
		query = slices.Compact(query)

		wordInDocs := getRelevance(query, index, queriesToShow)

		for i := range wordInDocs {
			fmt.Fprint(writer, wordInDocs[i].DocPosition, " ")
		}
		fmt.Fprintln(writer)
	}
}

type WordInDocs struct {
	DocPosition int
	WordPoints  int
}

// index := map[string][]WordInDocs
// key - слово
// value: список WordInDocs, где
// DocPosition - индекс документа (с единицы);
// Points - сколько раз слово встречается
func makeMap(docs [][]string) map[string][]WordInDocs {
	index := make(map[string][]WordInDocs)

	for i, doc := range docs {
		for _, word := range doc {
			sl := index[word]

			if len(sl) > 0 && sl[len(sl)-1].DocPosition == i+1 {
				sl[len(sl)-1].WordPoints++
				index[word] = sl
			} else {
				index[word] = append(index[word], WordInDocs{
					DocPosition: i + 1,
					WordPoints:  1,
				})
			}
		}
	}

	return index
}

func getTopN(queriesToShow int, relevance map[int]int) []WordInDocs {
	limit := min(len(relevance), queriesToShow)
	top := []WordInDocs{}

	for position, points := range relevance {
		elem := WordInDocs{
			DocPosition: position,
			WordPoints:  points,
		}
		top = collectTop(top, limit, elem)
	}

	return top
}

func getRelevance(query []string, index map[string][]WordInDocs,
	limit int) []WordInDocs {

	relevance := make(map[int]int)

	for _, word := range query {
		if wordInDocs, ok := index[word]; ok {
			for _, elem := range wordInDocs {
				relevance[elem.DocPosition] += elem.WordPoints
			}
		}
	}

	top := getTopN(limit, relevance)

	return top
}

// После выхода из collectTop слайс всегда отсортирован и содержит не более limit элементов
func collectTop(relevance []WordInDocs, limit int, elem WordInDocs) []WordInDocs {
	// slice is empty
	if len(relevance) == 0 {
		return append(relevance, elem)
	}

	// finding place to insert
	for i := range relevance {

		// new elem better than in relevance
		if elem.WordPoints > relevance[i].WordPoints ||
			(elem.WordPoints == relevance[i].WordPoints &&
				elem.DocPosition < relevance[i].DocPosition) {

			relevance = append(relevance, WordInDocs{})
			copy(relevance[i+1:], relevance[i:])
			relevance[i] = elem

			if len(relevance) > limit {
				relevance = relevance[:limit]
			}
			return relevance
		}

	}
	if len(relevance) < limit {
		relevance = append(relevance, elem)
	}
	return relevance
}
