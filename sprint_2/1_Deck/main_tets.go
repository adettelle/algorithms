package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		capacity int
		commands []string
	}{
		{
			"Test 1",
			7,
			10,
			[]string{
				"push_front -855",
				"push_front 0",
				"pop_back",
				"pop_back",
				"push_back 844",
				"pop_back",
				"push_back 823",
			},
		},
	}

	for _, tt := range tests {
		for _, command := range tt.commands {
			deck := NewDeck(tt.capacity)

			cmd := strings.Split(command, " ")
			val := 0
			if len(cmd) > 1 {
				val, _ = strconv.Atoi(cmd[1])
			}

			switch cmd[0] {
			case "pop_back":
				deck.PopBack()
			case "pop_front":
				deck.PopFront()
			case "push_front":
				deck.PushFront(val)
			case "push_back":
				deck.PushBack(val)
			}
		}

	}
}

/*
Ввод:
7
10
push_front -855
push_front 0
pop_back
pop_back
push_back 844
pop_back
push_back 823

Вывод:
-855
0
844

push_front -855: [-855 0 0 0 0 0 0 0 0 0] d.tail: 1 d.head: 0 size: 1
push_front Z:    [-855 0 0 0 0 0 0 0 0 Z] d.tail: 1 d.head: 9 size: 2
PopBack: -855    [0 0 0 0 0 0 0 0 0 Z]    d.tail: 0 d.head: 9 size: 1
PopBack: Z       [0 0 0 0 0 0 0 0 0 0]    d.tail: 0 d.head: 0 size: 0
push_back 844:   [844 0 0 0 0 0 0 0 0 0]  d.tail: 1 d.head: 0 size: 1
pop_back:		 [0 0 0 0 0 0 0 0 0 0]    d.tail: 0 d.head: 0 size: 0
push_back 823:   [823 0 0 0 0 0 0 0 0 0]  d.tail: 1 d.head: 0 size: 1
*/

/*
4
4
push_back 11
push_back 22
pop_front
push_back 33

Вывод: 11

PushBack 11: [11 0 0 0]  d.tail: 1 d.head: 0 d.size: 1
PushBack 22: [11 22 0 0] d.tail: 2 d.head: 0 d.size: 2
PopFront: 11 [0 22 0 0]  d.tail: 2 d.head: 1 d.size: 1
PushBack 33: [0 22 33 0] d.tail: 3 d.head: 1 d.size: 2
*/

/*
10
8
push_front 11
push_front 22
push_back 33
push_back 44
push_front 55
push_front 66
pop_back
pop_back
push_back 77
pop_front

Output: 44 33 66

PushFront 11: [11 0 0 0 0 0 0 0]      d.tail: 1 d.head: 0 d.size: 1
PushFront 22: [11 0 0 0 0 0 0 22]     d.tail: 1 d.head: 7 d.size: 2
PushBack 33:  [11 33 0 0 0 0 0 22]    d.tail: 2 d.head: 7 d.size: 3
PushBack 44:  [11 33 44 0 0 0 0 22]   d.tail: 3 d.head: 7 d.size: 4
PushFront 55: [11 33 44 0 0 0 55 22]  d.tail: 3 d.head: 6 d.size: 5
PushFront 66: [11 33 44 0 0 66 55 22] d.tail: 3 d.head: 5 d.size: 6
PopBack: 44   [11 33 0 0 0 66 55 22]  d.tail: 2 d.head: 5 d.size: 5
PopBack: 33   [11 0 0 0 0 66 55 22]   d.tail: 1 d.head: 5 d.size: 4
PushBack 77:  [11 77 0 0 0 66 55 22]  d.tail: 2 d.head: 5 d.size: 5
PopFront: 66  [11 77 0 0 0 0 55 22]   d.tail: 2 d.head: 6 d.size: 4
*/

/*
4
4
push_front 861
push_front -819
pop_back
pop_back

PushFront 861:  [861 0 0 0]    d.tail: 1 d.head: 0 d.size: 1
PushFront -819: [861 0 0 -819] d.tail: 1 d.head: 3 d.size: 2
PopBack: 861    [0 0 0 -819]   d.tail: 0 d.head: 3 d.size: 1
PopBack: -819   [0 0 0 0]      d.tail: 0 d.head: 0 d.size: 0
*/

/*
Ввод:
6
6
push_front -201
push_back 959
push_back 102
push_front 20
pop_front
pop_back

Вывод:
20
102

PushFront -201: [-201 0 0 0 0 0]      d.tail: 1 d.head: 0 d.size: 1
PushBack 959:   [-201 959 0 0 0 0]    d.tail: 2 d.head: 0 d.size: 2
PushBack 102:   [-201 959 102 0 0 0]  d.tail: 3 d.head: 0 d.size: 3
PushFront 20:   [-201 959 102 0 0 20] d.tail: 3 d.head: 5 d.size: 4
PopFront: 20    [-201 959 102 0 0 0]  d.tail: 3 d.head: 0 d.size: 3
PopBack: 102    [-201 959 0 0 0 0]    d.tail: 2 d.head: 0 d.size: 2
*/

/*
6
3
push_front 556
pop_front
push_front 229
pop_front
push_back -784
push_back -258

556
229

PushFront 556: [556 0 0]     d.tail: 1 d.head: 0 d.size: 1
PopFront: 556  [0 0 0]       d.tail: 0 d.head: 0 d.size: 0
PushFront:     [229 0 0]     d.tail: 1 d.head: 0 d.size: 1
PopFront: 229  [0 0 0]       d.tail: 0 d.head: 0 d.size: 0
PushBack -784: [-784 0 0]    d.tail: 1 d.head: 0 d.size: 1
PushBack -258: [-784 -258 0] d.tail: 2 d.head: 0 d.size: 2
*/
