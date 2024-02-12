package internal

import "fmt"

// Интерфейс состояния посылки
type ParcelState interface {
	ProcessParcel() string
}

// Конкретное состояние - посылка оформлена
type ParcelPlacedState struct{}

func (pps *ParcelPlacedState) ProcessParcel() string {
	return "Parcel is placed"
}

// Конкретное состояние - посылка отправлена
type ParcelShippedState struct{}

func (pss *ParcelShippedState) ProcessParcel() string {
	return "Parcel is shipped"
}

// Контекст - посылка
type ParcelContext struct {
	state ParcelState
}

func (pc *ParcelContext) SetState(state ParcelState) {
	pc.state = state
}

func (pc *ParcelContext) ProcessParcel() string {
	return pc.state.ProcessParcel()
}

// StatePattern точка входа
func StatePattern() {
	// Создаем посылку и устанавливаем начальное состояние (посылка оформлена)
	parcel := &ParcelContext{state: &ParcelPlacedState{}}

	// Обрабатываем посылку в текущем состоянии
	fmt.Println(parcel.ProcessParcel())

	// Меняем состояние на "посылка отправлена"
	parcel.SetState(&ParcelShippedState{})

	// Обрабатываем посылку в новом состоянии
	fmt.Println(parcel.ProcessParcel())
}
