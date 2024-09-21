package main

import (
	"fmt"
)

// Определяем структуру для мин-кучи
type MinHeap struct {
	elements []int // Срез для хранения элементов кучи
}

// Метод для добавления элемента в кучу
func (h *MinHeap) Insert(value int) {
	h.elements = append(h.elements, value) // Добавляем элемент в конец
	h.bubbleUp(len(h.elements) - 1)        // Восстанавливаем свойства кучи
}

// Метод для удаления минимального элемента (корня)
func (h *MinHeap) ExtractMin() int {
	if len(h.elements) == 0 {
		panic("Heap is empty") // Если куча пуста, вызываем панику
	}
	min := h.elements[0] // Минимальный элемент — корень
	lastIndex := len(h.elements) - 1
	h.elements[0] = h.elements[lastIndex] // Перемещаем последний элемент на место корня
	h.elements = h.elements[:lastIndex]   // Удаляем последний элемент
	h.bubbleDown(0)                       // Восстанавливаем свойства кучи
	return min
}

// Вспомогательный метод для восхождения элемента
func (h *MinHeap) bubbleUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.elements[index] >= h.elements[parentIndex] {
			return // Если текущий элемент больше или равен родителю, выходим
		}
		// Меняем местами текущий элемент и его родителя
		h.elements[index], h.elements[parentIndex] = h.elements[parentIndex], h.elements[index]
		index = parentIndex // Переходим к родителю
	}
}

// Вспомогательный метод для нисхождения элемента
func (h *MinHeap) bubbleDown(index int) {
	lastIndex := len(h.elements) - 1
	for index < lastIndex {
		leftChildIndex := 2*index + 1
		rightChildIndex := 2*index + 2
		smallerChildIndex := leftChildIndex

		if rightChildIndex <= lastIndex && h.elements[rightChildIndex] < h.elements[leftChildIndex] {
			smallerChildIndex = rightChildIndex // Выбираем меньший из двух дочерних узлов
		}

		if smallerChildIndex > lastIndex || h.elements[index] <= h.elements[smallerChildIndex] {
			return // Если текущий элемент меньше или равен меньшему дочернему, выходим
		}

		// Меняем местами текущий элемент и меньший дочерний
		h.elements[index], h.elements[smallerChildIndex] = h.elements[smallerChildIndex], h.elements[index]
		index = smallerChildIndex // Переходим к меньшему дочернему
	}
}

// Метод для отображения элементов кучи
func (h *MinHeap) Display() {
	fmt.Println(h.elements) // Выводим все элементы кучи
}

func main() {
	// Создаем новый экземпляр мин-кучи
	heap := &MinHeap{}

	// Добавляем элементы в кучу
	heap.Insert(5)
	heap.Insert(3)
	heap.Insert(8)
	heap.Insert(1)
	heap.Insert(4)

	fmt.Println("Элементы мин-кучи:")
	heap.Display() // Вывод: [1 3 8 5 4]

	// Извлекаем минимальный элемент
	min := heap.ExtractMin()
	fmt.Println("Извлеченный минимальный элемент:", min) // Вывод: 1

	fmt.Println("Элементы мин-кучи после извлечения минимального элемента:")
	heap.Display() // Вывод: [3 4 8 5]
}
