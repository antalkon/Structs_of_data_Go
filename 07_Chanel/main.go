package main

import (
	"fmt"
	"time"
)

// Функция, которая отправляет числа в канал
func sendNumbers(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Отправляю:", i)
		ch <- i                 // Отправка числа в канал
		time.Sleep(time.Second) // Пауза для имитации работы
	}
	close(ch) // Закрываем канал после отправки всех чисел
}

// Функция, которая получает числа из канала
func receiveNumbers(ch <-chan int) {
	for num := range ch { // Читаем из канала, пока он не закрыт
		fmt.Println("Получено:", num)
	}
}

func main() {
	// Создаем канал для передачи целых чисел
	ch := make(chan int)

	// Запускаем горутину для отправки чисел
	go sendNumbers(ch)

	// Запускаем горутину для получения чисел
	go receiveNumbers(ch)

	// Ждем немного, чтобы горутины успели завершить работу
	time.Sleep(6 * time.Second)
}
