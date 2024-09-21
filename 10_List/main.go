package main

import "fmt"

// Определяем структуру элемента односвязного списка
type Node struct {
	Value int   // Значение узла
	Next  *Node // Указатель на следующий узел
}

// Определяем структуру односвязного списка
type SinglyLinkedList struct {
	Head *Node // Указатель на первый элемент списка
}

// Метод для добавления элемента в конец списка
func (list *SinglyLinkedList) Append(value int) {
	newNode := &Node{Value: value} // Создаем новый узел

	if list.Head == nil { // Если список пуст
		list.Head = newNode // Новый узел становится головой списка
	} else {
		current := list.Head
		// Проходим до последнего узла
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode // Устанавливаем следующий узел
	}
}

// Метод для отображения элементов списка
func (list *SinglyLinkedList) Display() {
	current := list.Head // Начинаем с головы списка
	for current != nil {
		fmt.Print(current.Value, " ") // Выводим значение узла
		current = current.Next        // Переходим к следующему узлу
	}
	fmt.Println()
}

func main() {
	// Создаем экземпляр односвязного списка
	list := SinglyLinkedList{}

	// Добавляем элементы в список
	list.Append(1)
	list.Append(2)
	list.Append(3)

	fmt.Println("Односвязный список:")
	list.Display() // Вывод: 1 2 3
}
