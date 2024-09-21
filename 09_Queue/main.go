package main

import (
	"fmt"
)

// Определяем структуру Queue
type Queue struct {
	items []int // Слайс для хранения элементов очереди
}

// Метод для добавления элемента в очередь
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item) // Добавляем элемент в конец слайса
}

// Метод для извлечения элемента из очереди
func (q *Queue) Dequeue() (int, bool) {
	if len(q.items) == 0 { // Проверяем, пуста ли очередь
		return 0, false // Если пуста, возвращаем 0 и false
	}
	item := q.items[0]    // Получаем первый элемент
	q.items = q.items[1:] // Удаляем первый элемент
	return item, true     // Возвращаем извлеченный элемент и true
}

// Метод для проверки, пуста ли очередь
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0 // Возвращаем true, если очередь пуста
}

// Метод для получения размера очереди
func (q *Queue) Size() int {
	return len(q.items) // Возвращаем количество элементов в очереди
}

func main() {
	// Создаем экземпляр очереди
	queue := Queue{}

	// Добавляем элементы в очередь
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println("Размер очереди:", queue.Size()) // Вывод: 3

	// Извлекаем элементы из очереди
	for !queue.IsEmpty() {
		item, _ := queue.Dequeue()
		fmt.Println("Извлеченный элемент:", item)
	}

	// Проверка на пустоту
	fmt.Println("Очередь пуста:", queue.IsEmpty()) // Вывод: true
}
