package internal

import (
	"fmt"
)

// Интерфейс для всех продуктов
type Product interface {
	GetName() string
}

// Продукт A
type ProductA struct{}

func (p *ProductA) GetName() string {
	return "Product A"
}

// Продукт B
type ProductB struct{}

func (p *ProductB) GetName() string {
	return "Product B"
}

// Фабрика для создания продуктов
type Factory interface {
	Create() Product
}

// Фабрика для создания продукта A
type FactoryA struct{}

func (f *FactoryA) Create() Product {
	return &ProductA{}
}

// Фабрика для создания продукта B
type FactoryB struct{}

func (f *FactoryB) Create() Product {
	return &ProductB{}
}

// Функция, которая принимает фабрику и возвращает продукт
func GetProduct(factory Factory) Product {
	return factory.Create()
}

func FactoryPattern() {
	// Создаем фабрику A
	factoryA := &FactoryA{}

	// Получаем продукт A от фабрики A
	productA := GetProduct(factoryA)
	fmt.Println("Product from factory A:", productA.GetName())

	// Создаем фабрику B
	factoryB := &FactoryB{}

	// Получаем продукт B от фабрики B
	productB := GetProduct(factoryB)
	fmt.Println("Product from factory B:", productB.GetName())
}
