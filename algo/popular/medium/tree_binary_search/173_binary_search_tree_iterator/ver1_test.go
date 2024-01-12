package main

import "testing"

func TestBSTIterator(t *testing.T) {
	// Создаем тестовый BST: 7, 3, 15, null, null, 9, 20
	root := &TreeNode{Val: 7}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 15}
	root.Right.Left = &TreeNode{Val: 9}
	root.Right.Right = &TreeNode{Val: 20}

	iterator := Constructor(root)

	// Проверяем, что HasNext() возвращает true, так как у нас есть элементы
	if !iterator.HasNext() {
		t.Errorf("Expected HasNext() to return true, but got false")
	}

	// Проверяем, что Next() возвращает наименьший элемент
	if val := iterator.Next(); val != 3 {
		t.Errorf("Expected Next() to return 3, but got %d", val)
	}

	// Проверяем, что HasNext() возвращает true, так как у нас еще есть элементы
	if !iterator.HasNext() {
		t.Errorf("Expected HasNext() to return true, but got false")
	}

	// Проверяем, что Next() возвращает следующий наименьший элемент
	if val := iterator.Next(); val != 7 {
		t.Errorf("Expected Next() to return 7, but got %d", val)
	}

	// Проверяем, что HasNext() возвращает true, так как у нас еще есть элементы
	if !iterator.HasNext() {
		t.Errorf("Expected HasNext() to return true, but got false")
	}

	// Проверяем, что Next() возвращает следующий наименьший элемент
	if val := iterator.Next(); val != 9 {
		t.Errorf("Expected Next() to return 9, but got %d", val)
	}

	// Проверяем, что HasNext() возвращает true, так как у нас еще есть элементы
	if !iterator.HasNext() {
		t.Errorf("Expected HasNext() to return true, but got false")
	}

	// Проверяем, что Next() возвращает следующий наименьший элемент
	if val := iterator.Next(); val != 15 {
		t.Errorf("Expected Next() to return 15, but got %d", val)
	}

	// Проверяем, что HasNext() возвращает true, так как у нас еще есть элементы
	if !iterator.HasNext() {
		t.Errorf("Expected HasNext() to return true, but got false")
	}

	// Проверяем, что Next() возвращает следующий наименьший элемент
	if val := iterator.Next(); val != 20 {
		t.Errorf("Expected Next() to return 20, but got %d", val)
	}

	// Проверяем, что HasNext() возвращает false, так как у нас больше нет элементов
	if iterator.HasNext() {
		t.Errorf("Expected HasNext() to return false, but got true")
	}
}
