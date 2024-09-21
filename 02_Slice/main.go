package main

import "fmt"

func main() {
	// Объявление и инициализация пустого среза целых чисел
	var numbers []int
	fmt.Println("Пустой срез:", numbers) // Вывод: []

	// Добавление элементов в срез с помощью функции append
	numbers = append(numbers, 10)
	numbers = append(numbers, 20, 30)
	fmt.Println("Срез после добавления элементов:", numbers) // Вывод: [10 20 30]

	// Объявление и инициализация среза с помощью литерала
	fruits := []string{"яблоко", "банан", "апельсин"}
	fmt.Println("Фрукты:", fruits) // Вывод: [яблоко банан апельсин]

	// Длина и емкость среза
	fmt.Println("Длина среза fruits:", len(fruits))   // Вывод: 3
	fmt.Println("Емкость среза fruits:", cap(fruits)) // Вывод: 3

	// Изменение элемента среза
	fruits[1] = "груша"
	fmt.Println("После изменения второго элемента:", fruits) // Вывод: [яблоко груша апельсин]

	// Создание среза с помощью make
	colors := make([]string, 0, 5) // Создаем пустой срез длиной 0 и емкостью 5
	colors = append(colors, "красный", "зеленый", "синий")
	fmt.Println("Цвета:", colors) // Вывод: [красный зеленый синий]

	// Подсрезы: создание нового среза из существующего
	subFruits := fruits[1:3]                   // Создаем срез, который включает элементы с 1 по 2
	fmt.Println("Подсрез фруктов:", subFruits) // Вывод: [груша апельсин]

	// Копирование среза с помощью copy
	copiedFruits := make([]string, len(fruits)) // Создаем новый срез
	copy(copiedFruits, fruits)                  // Копируем элементы
	fmt.Println("Копия фруктов:", copiedFruits) // Вывод: [яблоко груша апельсин]
}
