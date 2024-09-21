package main

import "fmt"

func main() {
	// Объявляем и инициализируем пустую карту
	phoneBook := make(map[string]string)

	// Добавляем записи в карту
	phoneBook["Иван"] = "123-456"
	phoneBook["Анна"] = "789-101"
	phoneBook["Сергей"] = "112-131"
	fmt.Println("Телефонная книга:", phoneBook)

	// Получение значения по ключу
	ivanPhone := phoneBook["Иван"]
	fmt.Println("Телефон Ивана:", ivanPhone) // Вывод: Телефон Ивана: 123-456

	// Проверка наличия ключа в карте
	if phone, exists := phoneBook["Анна"]; exists {
		fmt.Println("Телефон Анны:", phone) // Вывод: Телефон Анны: 789-101
	} else {
		fmt.Println("Анна не найдена в телефонной книге.")
	}

	// Удаление записи из карты
	delete(phoneBook, "Сергей")
	fmt.Println("После удаления Сергея:", phoneBook)

	// Перебор элементов карты
	fmt.Println("Все записи в телефонной книге:")
	for name, phone := range phoneBook {
		fmt.Printf("Имя: %s, Телефон: %s\n", name, phone)
	}

	// Инициализация карты с данными
	studentGrades := map[string]int{
		"Иван":   90,
		"Анна":   85,
		"Сергей": 92,
	}
	fmt.Println("Оценки студентов:", studentGrades)
}
