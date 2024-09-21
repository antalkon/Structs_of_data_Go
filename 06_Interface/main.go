package main

import "fmt"

// Объявляем интерфейс Animal с методом Speak
type Animal interface {
	Speak() string
}

// Объявляем структуру Dog, реализующую интерфейс Animal
type Dog struct {
	Name string
}

// Метод Speak для структуры Dog
func (d Dog) Speak() string {
	return "Гав! Я " + d.Name
}

// Объявляем структуру Cat, также реализующую интерфейс Animal
type Cat struct {
	Name string
}

// Метод Speak для структуры Cat
func (c Cat) Speak() string {
	return "Мяу! Я " + c.Name
}

// Функция, принимающая интерфейс Animal
func makeSound(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	// Создаем экземпляры Dog и Cat
	dog := Dog{Name: "Бобик"}
	cat := Cat{Name: "Мурка"}

	// Вызываем функцию makeSound с разными типами
	makeSound(dog) // Вывод: Гав! Я Бобик
	makeSound(cat) // Вывод: Мяу! Я Мурка
}
