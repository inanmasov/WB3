package internal

import "fmt"

// Интерфейс посещения ресторанов
type Restaurant interface {
	Accept(Visitor)
}

// Ресторан А
type RestaurantA struct {
}

func (r *RestaurantA) Accept(v Visitor) {
	v.VisitRestaurantA()
}

// Ресторан B
type RestaurantB struct {
}

func (r *RestaurantB) Accept(v Visitor) {
	v.VisitRestaurantB()
}

// Интерфейс посещения ресторанов
type Visitor interface {
	VisitRestaurantA()
	VisitRestaurantB()
}

// Врач, который может посетить 2 ресторана
type Doctor struct {
}

func (d *Doctor) VisitRestaurantA() {
	fmt.Println("The doctor visited the restaurant A")
}

func (d *Doctor) VisitRestaurantB() {
	fmt.Println("The doctor visited the restaurant B")
}

// Юрист, который может посетить 2 ресторана
type Lawyer struct {
}

func (l *Lawyer) VisitRestaurantA() {
	fmt.Println("The lawyer visited the restaurant A")

}

func (l *Lawyer) VisitRestaurantB() {
	fmt.Println("The lawyer visited the restaurant B")
}

func VisitorPattern() {
	// Список ресторанов, которые можно посетить
	restaurants := []Restaurant{&RestaurantA{}, &RestaurantB{}}
	// Gjctnbntkb
	doctor := &Doctor{}
	lawyer := &Lawyer{}
	// Посещаем рестораны
	for _, restaurant := range restaurants {
		restaurant.Accept(doctor)
		restaurant.Accept(lawyer)
	}
}
