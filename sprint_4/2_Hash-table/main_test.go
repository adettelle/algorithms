package main

import "testing"

func TestPut(t *testing.T) {
	tests := []struct {
		name          string
		initial       []Pair
		key           int
		value         int
		expectedValue int
		expectedCount int
	}{
		{
			name:          "insert new key",
			initial:       nil,
			key:           1,
			value:         100,
			expectedValue: 100,
			expectedCount: 1,
		},
		{
			name: "update existing key",
			initial: []Pair{
				{Key: 1, Value: 100},
			},
			key:           1,
			value:         200,
			expectedValue: 200,
			expectedCount: 1,
		},
		{
			name: "insert second key",
			initial: []Pair{
				{Key: 1, Value: 100},
			},
			key:           2,
			value:         300,
			expectedValue: 300,
			expectedCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash := NewHashTable(10)

			// Заполняем таблицу начальными данными.
			for _, p := range tt.initial {
				hash.put(p.Key, p.Value)
			}

			// Тестируем put.
			hash.put(tt.key, tt.value)

			// Проверяем значение.
			got, ok := hash.get(tt.key)
			if !ok {
				t.Errorf("key %d not found", tt.key)
			}

			if got != tt.expectedValue {
				t.Errorf("expected value %d, got %d", tt.expectedValue, got)
			}

			// Проверяем количество элементов.
			count := 0
			for _, bucket := range hash.buckets {
				count += len(bucket)
			}

			if count != tt.expectedCount {
				t.Errorf("expected %d elements, got %d", tt.expectedCount, count)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name          string
		initial       []Pair
		key           int
		expectedValue int
		expectedCount int
		expectedOK    bool
	}{
		{
			name:          "get unexisting key",
			initial:       nil,
			key:           1,
			expectedValue: 0,
			expectedCount: 0,
			expectedOK:    false,
		},
		{
			name: "get existing key",
			initial: []Pair{
				{Key: 1, Value: 100},
			},
			key:           1,
			expectedValue: 100,
			expectedCount: 1,
			expectedOK:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash := NewHashTable(10)

			// Заполняем таблицу начальными данными.
			for _, p := range tt.initial {
				hash.put(p.Key, p.Value)
			}

			// Тестируем get.
			got, ok := hash.get(tt.key)
			if ok != tt.expectedOK {
				t.Errorf("key found: expected %v, result %v", tt.expectedOK, ok)
			}

			if got != tt.expectedValue {
				t.Errorf("expected value %d, got %d", tt.expectedValue, got)
			}

			// Проверяем количество элементов.
			count := 0
			for _, bucket := range hash.buckets {
				count += len(bucket)
			}

			if count != tt.expectedCount {
				t.Errorf("expected %d elements, got %d", tt.expectedCount, count)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name          string
		initial       []Pair
		key           int
		value         int
		keyToDelete   int
		expectedValue int
		expectedCount int
		expectedOK    bool
	}{
		{
			name: "delete existing key",
			initial: []Pair{
				{Key: 1, Value: 100},
			},
			key:           1,
			value:         100,
			keyToDelete:   1,
			expectedValue: 100,
			expectedCount: 0,
			expectedOK:    true,
		},
		{
			name: "delete unexisting key",
			initial: []Pair{
				{Key: 1, Value: 100},
			},
			key:           1,
			value:         100,
			keyToDelete:   2,
			expectedValue: 0,
			expectedCount: 1,
			expectedOK:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash := NewHashTable(10)

			// Заполняем таблицу начальными данными.
			for _, p := range tt.initial {
				hash.put(p.Key, p.Value)
			}

			// Тестируем delete.
			got, ok := hash.delete(tt.keyToDelete)
			if ok != tt.expectedOK {
				t.Errorf("key found: expected %v, result %v", tt.expectedOK, ok)
			}

			if got != tt.expectedValue {
				t.Errorf("expected value %d, got %d", tt.expectedValue, got)
			}

			// Проверяем количество элементов.
			count := 0
			for _, bucket := range hash.buckets {
				count += len(bucket)
			}

			if count != tt.expectedCount {
				t.Errorf("expected %d elements, got %d", tt.expectedCount, count)
			}
		})
	}
}
