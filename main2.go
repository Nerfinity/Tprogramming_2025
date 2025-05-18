package main

import (
    "fmt"
    
)

// Структура Dish
type Dish struct {
    Name        string
    Price       float64
    Description string
}

// Метод NewDish - конструктор, возвращающий структуру Dish
func NewDish(name string, price float64, description string) *Dish {
    return &Dish{
        Name:        name,
        Price:       price,
        Description: description,
    }
}

// Метод GetInfo возвращает информацию о блюде
func (d *Dish) GetInfo() string {
    return fmt.Sprintf("Блюдо: %s, Цена: %.2f, Описание: %s", d.Name, d.Price, d.Description)
}

// Метод SetPrice изменяет цену блюда
func (d *Dish) SetPrice(newPrice float64) {
    d.Price = newPrice
}

// Метод SetDescription изменяет описание блюда
func (d *Dish) SetDescription(newDescription string) {
    d.Description = newDescription
}

func main() {
    // Создаем несколько блюд с помощью конструктора
    dish1 := NewDish("Паста Карбонара", 15.99, "Классическая итальянская паста с беконом и яйцом")
    dish2 := NewDish("Салат Цезарь", 12.99, "Салат с курицей, сыром и гренками")
    dish3 := NewDish("Вегетарианская пицца", 10.99, "Пицца с овощами и сыром")

    // Выводим информацию о каждом блюде
    fmt.Println(dish1.GetInfo())
    fmt.Println(dish2.GetInfo())
    fmt.Println(dish3.GetInfo())

    // Изменяем цену блюда
    dish1.SetPrice(18.99)

    // Изменяем описание блюда
    dish2.SetDescription("Салат с курицей, сыром и гренками, с добавлением помидоров")

    // Выводим обновленную информацию о каждом блюде
    fmt.Println(dish1.GetInfo())
    fmt.Println(dish2.GetInfo())
    fmt.Println(dish3.GetInfo())
}