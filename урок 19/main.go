package main

import (
	"fmt"
	"math"
)

type Person struct {
	Name string
	Age  int
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Student struct {
	Name  string
	Grade int
}

type Circle struct {
	Radius float64
}

// Структура для задания 7.
type Car struct {
	Brand string
	Speed int
}

func CheckSpeed(car Car) string {
	if car.Speed > 100 {
		return "Слишком быстро"
	}
	if car.Speed >= 60 {
		return "Нормальная скорость"
	}
	return "Медленно"
}

func main() {
	// Задание 1
	fmt.Println("Задание 1")
	person := Person{Name: "Алексей", Age: 20}
	fmt.Printf("Имя: %s, возраст: %d\n", person.Name, person.Age)

	// Задание 2
	fmt.Println("\nЗадание 2")
	book := Book{
		Title:  "Капитанская дочка",
		Author: "Александр Пушкин",
		Pages:  200,
	}
	fmt.Printf("Название: %s\nАвтор: %s\nСтраниц: %d\n",
		book.Title, book.Author, book.Pages)

	// Задание 3
	fmt.Println("\nЗадание 3")
	{
		type Car struct {
			Brand string
			Year  int
		}

		car := Car{Brand: "Toyota", Year: 2015}
		carPointer := &car
		carPointer.Year = 2020

		fmt.Printf("Марка: %s, год: %d\n", car.Brand, car.Year)
	}

	// Задание 4
	fmt.Println("\nЗадание 4")
	rectangle := Rectangle{Width: 5, Height: 3}
	fmt.Println("Площадь:", rectangle.Width*rectangle.Height)
	fmt.Println("Периметр:", 2*(rectangle.Width+rectangle.Height))

	rectangle.Width = 8
	rectangle.Height = 4
	fmt.Println("Новая площадь:", rectangle.Width*rectangle.Height)
	fmt.Println("Новый периметр:", 2*(rectangle.Width+rectangle.Height))

	// Задание 5
	fmt.Println("\nЗадание 5")
	student1 := Student{Name: "Анна", Grade: 85}
	student2 := Student{Name: "Иван", Grade: 92}

	fmt.Printf("Студент %s получил %d баллов\n",
		student1.Name, student1.Grade)
	fmt.Printf("Студент %s получил %d баллов\n",
		student2.Name, student2.Grade)

	if student1.Grade > student2.Grade {
		fmt.Println("Больше баллов получил:", student1.Name)
	} else if student2.Grade > student1.Grade {
		fmt.Println("Больше баллов получил:", student2.Name)
	} else {
		fmt.Println("Студенты получили одинаковое количество баллов")
	}

	// Задание 6
	fmt.Println("\nЗадание 6")
	circle := Circle{Radius: 3}
	fmt.Printf("Площадь: %.2f\n", math.Pi*circle.Radius*circle.Radius)
	fmt.Printf("Длина окружности: %.2f\n", 2*math.Pi*circle.Radius)

	circle.Radius = 5
	fmt.Printf("Новая площадь: %.2f\n", math.Pi*circle.Radius*circle.Radius)
	fmt.Printf("Новая длина окружности: %.2f\n", 2*math.Pi*circle.Radius)

	// Задание 7
	fmt.Println("\nЗадание 7")
	car1 := Car{Brand: "BMW", Speed: 120}
	car2 := Car{Brand: "Toyota", Speed: 80}
	car3 := Car{Brand: "Lada", Speed: 40}

	fmt.Printf("Машина %s: %s\n", car1.Brand, CheckSpeed(car1))
	fmt.Printf("Машина %s: %s\n", car2.Brand, CheckSpeed(car2))
	fmt.Printf("Машина %s: %s\n", car3.Brand, CheckSpeed(car3))
}
