package main

import (
	"fmt"
)

// Определяем структуру Stack
type Stack struct {
	items []int // Слайс для хранения элементов стека
}

// Метод для добавления элемента в стек
func (s *Stack) Push(item int) {
	s.items = append(s.items, item) // Добавляем элемент в конец слайса
}

// Метод для извлечения элемента из стека
func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 { // Проверяем, пуст ли стек
		return 0, false // Если пуст, возвращаем 0 и false
	}
	// Получаем последний элемент
	index := len(s.items) - 1
	item := s.items[index]
	// Удаляем последний элемент
	s.items = s.items[:index]
	return item, true // Возвращаем извлеченный элемент и true
}

// Метод для проверки, пуст ли стек
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0 // Возвращаем true, если стек пуст
}

// Метод для получения размера стека
func (s *Stack) Size() int {
	return len(s.items) // Возвращаем количество элементов в стеке
}

func main() {
	// Создаем экземпляр стека
	stack := Stack{}

	// Добавляем элементы в стек
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("Размер стека:", stack.Size()) // Вывод: 3

	// Извлекаем элементы из стека
	for !stack.IsEmpty() {
		item, _ := stack.Pop()
		fmt.Println("Извлеченный элемент:", item)
	}

	// Проверка на пустоту
	fmt.Println("Стек пуст:", stack.IsEmpty()) // Вывод: true
}
