package main

import (
	"fmt"
)

// Определяем структуру для графа
type Graph struct {
	vertices map[int][]int // Хранит список смежности, где ключ — вершина, а значение — срез смежных вершин
}

// Создаем новый граф
func NewGraph() *Graph {
	return &Graph{vertices: make(map[int][]int)} // Инициализируем граф с пустым списком смежности
}

// Метод для добавления вершины
func (g *Graph) AddVertex(vertex int) {
	if _, exists := g.vertices[vertex]; !exists {
		g.vertices[vertex] = []int{} // Создаем запись для новой вершины
	}
}

// Метод для добавления ребра
func (g *Graph) AddEdge(v1, v2 int) {
	g.AddVertex(v1)                             // Убедимся, что вершина v1 существует
	g.AddVertex(v2)                             // Убедимся, что вершина v2 существует
	g.vertices[v1] = append(g.vertices[v1], v2) // Добавляем v2 в список смежных вершин для v1
	g.vertices[v2] = append(g.vertices[v2], v1) // Для неориентированного графа добавляем v1 в список v2
}

// Метод для отображения графа
func (g *Graph) Display() {
	for vertex, edges := range g.vertices {
		fmt.Printf("%d: %v\n", vertex, edges) // Выводим вершину и её смежные вершины
	}
}

func main() {
	// Создаем новый граф
	graph := NewGraph()

	// Добавляем ребра в граф
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 4)

	fmt.Println("Список смежности графа:")
	graph.Display()
}
