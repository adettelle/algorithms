package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	arr := make([]Lesson, n)
	for i := 0; i < n; i++ {
		var s, e string
		fmt.Fscan(reader, &s, &e)
		start := readInput(s)
		end := readInput(e)
		lesson := Lesson{start: start, end: end}
		arr[i] = lesson
	}

	sort.Slice(arr, func(i, j int) bool {
		return Less(arr[i], arr[j])
	})

	res := []Lesson{arr[0]}

	for i := 1; i < len(arr); i++ {
		res = checkElem(res, arr[i])
	}
	count := len(res)

	fmt.Println(count)

	sort.Slice(res, func(i, j int) bool {
		return LessOut(res[i], res[j])
	})

	for i := 0; i < count; i++ {
		start := getOutput(res[i].start)
		end := getOutput(res[i].end)
		fmt.Println(start, end)
	}
}

type Lesson struct {
	start int
	end   int
}

func readInput(timeInput string) int {
	minutesInH := 60
	before, after, ok := strings.Cut(timeInput, ".")
	if !ok {
		h, _ := strconv.Atoi(before)
		return h * minutesInH
	}

	hours, _ := strconv.Atoi(before)
	minutes, _ := strconv.Atoi(after)
	if len(after) == 1 {
		minutes = minutes * 10
	}

	return hours*minutesInH + minutes
}

func getOutput(minutes int) string {
	minutesInH := 60
	h := minutes / minutesInH
	m := minutes % minutesInH

	// fmt.Println(h, m)
	if m == 0 {
		// fmt.Println(strconv.Itoa(h))
		return strconv.Itoa(h)
	}

	hh := strconv.Itoa(h)
	mm := strconv.Itoa(m)

	// fmt.Println(strings.Join([]string{hh, mm}, "."))
	return strings.Join([]string{hh, mm}, ".")
}

// сортировка по длине урока, затем по времени начала, далее по времени конца урока
func Less(lesson1, lesson2 Lesson) bool {
	duration1 := lesson1.end - lesson1.start
	duration2 := lesson2.end - lesson2.start

	if duration1 != duration2 {
		return duration1 < duration2
	}

	return lesson1.start < lesson2.start
}

// сортировка по времени начала урока, далее по времени конца урока
func LessOut(lesson1, lesson2 Lesson) bool {
	if lesson1.start != lesson2.start {
		return lesson1.start < lesson2.start
	}
	return lesson1.end < lesson2.end
}

func checkElem(res []Lesson, elem Lesson) []Lesson {
	for i := 0; i < len(res); i++ {
		if res[i].end <= elem.start || res[i].start >= elem.end {
			continue
		} else {
			return res
		}
	}
	res = append(res, elem)
	return res
}
