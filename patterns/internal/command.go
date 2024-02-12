package internal

import "fmt"

// Интерфейс команды
type Command interface {
	OpenDoor()
}

// Конкретная команда
type OpeningDoors struct {
	door *Door
}

func (od *OpeningDoors) OpenDoor() {
	od.door.Action()
}

// Интерфейс получателя
type Door struct{}

func (r *Door) Action() {
	fmt.Println("Открытие дввери")
}

// Пульт интернета вещей - хранит команду и вызывает ее
type ControllerIoT struct {
	command Command
}

func (i *ControllerIoT) SetCommand(command Command) {
	i.command = command
}

func (i *ControllerIoT) OpenDoor() {
	i.command.OpenDoor()
}

func CommandPattern() {
	command := &OpeningDoors{door: &Door{}}
	invoker := &ControllerIoT{command: command}

	// Выполняем команду через инвокера
	invoker.OpenDoor()
}
