package internal

import "fmt"

// Данные
type Data struct {
	text         string
	confidential bool
}

type Handler interface {
	HandleRequest(data Data)
	SetNextHandler(nextHandler Handler)
}

// Конкретный обработчик - Руководитель отдела
type DepartmentManager struct {
	nextHandler Handler
}

func (dm *DepartmentManager) HandleRequest(data Data) {
	if !data.confidential {
		fmt.Printf("Неконфиденциальные данные \"%s\" отправлены руководителю отдела\n", data.text)
	} else {
		fmt.Println("Конфиденциальные данные отправлены вышестоящемоу руководству")
		dm.nextHandler.HandleRequest(data)
	}
}

func (dm *DepartmentManager) SetNextHandler(nextHandler Handler) {
	dm.nextHandler = nextHandler
}

// Конкретный обработчик - Генеральный директор
type CEO struct {
	nextHandler Handler
}

func (ceo *CEO) HandleRequest(data Data) {
	if data.confidential {
		fmt.Printf("Конфиденциальные данные \"%s\" отправлены генеральному директору\n", data.text)
	} else { // Впоследствии можно добавлять логику
		if ceo.nextHandler != nil {
			fmt.Println("Данные отправлены вышестоящемоу руководству")
			ceo.nextHandler.HandleRequest(data)
		} else {
			fmt.Println("Данные дошли до самых высоких должнотстей")
		}
	}
}

func (ceo *CEO) SetNextHandler(nextHandler Handler) {
	ceo.nextHandler = nextHandler
}

func ChainOfResponsibilityPattern() {
	// Создаем обработчиков
	departmentManager := &DepartmentManager{}
	ceo := &CEO{}

	// Настраиваем цепочку обработчиков
	departmentManager.SetNextHandler(ceo)

	// Создаем запросы на отпуск
	data1 := Data{text: "John", confidential: false}
	data2 := Data{text: "ifmd&%#0jfn$", confidential: true}

	// Обрабатываем запросы с помощью цепочки обработчиков
	departmentManager.HandleRequest(data1)
	departmentManager.HandleRequest(data2)
}
