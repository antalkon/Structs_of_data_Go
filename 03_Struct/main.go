package main

import "fmt"

// Объявляем структуру для представления человека
type Person struct {
	Name    string // Имя человека
	Age     int    // Возраст человека
	Address string // Адрес человека
}

func main() {
	// Создаем экземпляр структуры Person с использованием литерала
	person1 := Person{
		Name:    "Иван",
		Age:     30,
		Address: "Москва, ул. Ленина, 1",
	}

	// Выводим значения полей структуры
	fmt.Println("Информация о человеке:", person1)

	// Объявляем другой экземпляр структуры с использованием нового синтаксиса
	person2 := Person{"Анна", 25, "Санкт-Петербург, ул. Пушкина, 2"}
	fmt.Println("Информация о человеке 2:", person2)

	// Изменение поля структуры
	person1.Age = 31
	fmt.Println("Обновленная информация о человеке:", person1)

	// Использование указателей на структуры
	updateAge(&person1, 32) // Обновляем возраст с помощью функции
	fmt.Println("После обновления возраста:", person1)

	// Вложенные структуры
	type Address struct {
		Street string
		City   string
	}

	type PersonWithAddress struct {
		Name    string
		Age     int
		Address Address // Вложенная структура
	}

	// Создаем экземпляр структуры с вложенной структурой
	person3 := PersonWithAddress{
		Name: "Сергей",
		Age:  28,
		Address: Address{
			Street: "ул. Невского, 3",
			City:   "Казань",
		},
	}
	fmt.Println("Информация о человеке с адресом:", person3)
}

// Функция для обновления возраста человека с помощью указателя
func updateAge(p *Person, newAge int) {
	p.Age = newAge
}
