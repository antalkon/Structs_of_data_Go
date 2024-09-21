package main

import "fmt"

// Определяем структуру элемента двусвязного списка
type DoublyNode struct {
	Value int         // Значение узла
	Prev  *DoublyNode // Указатель на предыдущий узел
	Next  *DoublyNode // Указатель на следующий узел
}

// Определяем структуру двусвязного списка
type DoublyLinkedList struct {
	Head *DoublyNode // Указатель на первый элемент списка
	Tail *DoublyNode // Указатель на последний элемент списка
}

// Метод для добавления элемента в конец списка
func (list *DoublyLinkedList) Append(value int) {
	newNode := &DoublyNode{Value: value} // Создаем новый узел

	if list.Head == nil { // Если список пуст
		list.Head = newNode // Новый узел становится головой списка
		list.Tail = newNode // И хвостом
	} else {
		newNode.Prev = list.Tail // Устанавливаем указатель на предыдущий узел
		list.Tail.Next = newNode // Устанавливаем указатель на новый узел
		list.Tail = newNode      // Обновляем хвост списка
	}
}

// Метод для отображения элементов списка
func (list *DoublyLinkedList) Display() {
	current := list.Head // Начинаем с головы списка
	for current != nil {
		fmt.Print(current.Value, " ") // Выводим значение узла
		current = current.Next        // Переходим к следующему узлу
	}
	fmt.Println()
}

func main() {
	// Создаем экземпляр двусвязного списка
	doublyList := DoublyLinkedList{}

	// Добавляем элементы в список
	doublyList.Append(1)
	doublyList.Append(2)
	doublyList.Append(3)

	fmt.Println("Двусвязный список:")
	doublyList.Display() // Вывод: 1 2 3
}
