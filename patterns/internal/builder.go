package internal

import "fmt"

// Машина
type Car struct {
	Color     string
	Engine    string
	SeatCount int
	Brand     string
}

// Интерфейс строителя
type Builder interface {
	BuildColor()
	BuildEngine()
	BuildSeatCount()
	BuildBrand()
	GetCar() Car
}

// Инженер, который назначает строителя
type Engineer struct {
	builder Builder
}

// Конструктор инженера
func NewEngineer(builder Builder) *Engineer {
	return &Engineer{
		builder: builder,
	}
}

// Создание машины
func (e *Engineer) BuildCar() Car {
	e.builder.BuildColor()
	e.builder.BuildBrand()
	e.builder.BuildEngine()
	e.builder.BuildSeatCount()
	return e.builder.GetCar()
}

// Инженер класса А
type EngineerClassA struct {
	car Car
}

func (a *EngineerClassA) BuildColor() {
	a.car.Color = "green"
}

func (a *EngineerClassA) BuildBrand() {
	a.car.Brand = "toyota"
}

func (a *EngineerClassA) BuildEngine() {
	a.car.Engine = "5.2"
}

func (a *EngineerClassA) BuildSeatCount() {
	a.car.SeatCount = 4
}

func (a *EngineerClassA) GetCar() Car {
	return a.car
}

// Инженер класса B
type EngineerClassB struct {
	car Car
}

func (b *EngineerClassB) BuildColor() {
	b.car.Color = "red"
}

func (b *EngineerClassB) BuildBrand() {
	b.car.Brand = "lamborghini"
}

func (b *EngineerClassB) BuildEngine() {
	b.car.Engine = "9.4"
}

func (b *EngineerClassB) BuildSeatCount() {
	b.car.SeatCount = 2
}

func (b *EngineerClassB) GetCar() Car {
	return b.car
}

func BuilderPattern() {
	classA := &EngineerClassA{}
	engineer := NewEngineer(classA)
	carClassA := engineer.BuildCar()
	fmt.Println("Car class A:")
	fmt.Println(carClassA)

	classB := &EngineerClassB{}
	engineer = NewEngineer(classB)
	carClassB := engineer.BuildCar()
	fmt.Println("Car class B:")
	fmt.Println(carClassB)
}
