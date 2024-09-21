package main

import "fmt"

// Определяем структуру узла дерева
type TreeNode struct {
	Value int       // Значение узла
	Left  *TreeNode // Указатель на левое поддерево
	Right *TreeNode // Указатель на правое поддерево
}

// Метод для добавления значения в бинарное дерево
func (node *TreeNode) Insert(value int) {
	if value < node.Value { // Если значение меньше текущего
		if node.Left == nil {
			node.Left = &TreeNode{Value: value} // Создаем новый узел
		} else {
			node.Left.Insert(value) // Рекурсивно вставляем в левое поддерево
		}
	} else { // Если значение больше или равно текущему
		if node.Right == nil {
			node.Right = &TreeNode{Value: value} // Создаем новый узел
		} else {
			node.Right.Insert(value) // Рекурсивно вставляем в правое поддерево
		}
	}
}

// Метод для обхода дерева в симметричном порядке (in-order)
func (node *TreeNode) InOrder() {
	if node != nil {
		node.Left.InOrder()        // Обходим левое поддерево
		fmt.Print(node.Value, " ") // Выводим значение узла
		node.Right.InOrder()       // Обходим правое поддерево
	}
}

func main() {
	// Создаем корень дерева
	root := &TreeNode{Value: 10}

	// Добавляем значения в дерево
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(18)

	fmt.Println("Симметричный обход дерева:")
	root.InOrder() // Вывод: 3 5 7 10 12 15 18
}
