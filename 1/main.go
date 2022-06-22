package main

import "fmt"

type Human struct {
	Name    string
	Address string
	Phone   string
	Email   string
}

func (h *Human) SetName(name string) {
	h.Name = name
}

func (h *Human) SetPhone(phone string) {
	h.Phone = phone
}

func (h *Human) SetAddress(address string) {
	h.Address = address
}

type Action struct {
	Name       string
	SecondName string
	Address    string
	Human
	H Human
}

func (a *Action) SetName(name string) {
	a.Name = name
}

func main() {
	var action Action = Action{
		Name:       "John",
		SecondName: "Doe",
		Address:    "qwerty",
		Human: Human{
			Name:    "Jane",
			Address: "456",
			Phone:   "+79991234578",
			Email:   "john@doe.com",
		},
		H: Human{
			Name:    "Peter",
			Address: "123",
			Phone:   "+74561237845",
			Email:   "peter@example.com",
		},
	}

	/*
		Обращение к полям
	*/

	// при наличии одикаковых полей, приоритет у поля более верхнего уровня
	fmt.Println(action.Name) // Jonn

	// чтобы обратиться к полю встроенной структуре, нужно указать ее тип, затем поле
	fmt.Println(action.Human.Name) // Jane

	// если поле уникальное, то можно напрямую обращаться к нему из родительской структуры
	fmt.Println(action.Address) // 456

	// чтобы обратиться к полю типа структуры, нужно указать имя поля, затем необходимое поле вложенной структуры
	fmt.Println(action.H.Name) // Peter

	/*
		Вызов методов
	*/

	// методы встроенной структуры наследуются родительской
	action.SetPhone("+712345")
	fmt.Println(action.Phone) // +712345

	// при наличии одинаковых методов, приоритет у метода структуры более верхнего уровня
	action.SetName("new name")
	fmt.Println(action.Name)       // new name
	fmt.Println(action.Human.Name) // Jane

	// чтобы вызвать этот метод у встроенной структуры, необходимо обратиться к ней через тип
	action.Human.SetName("Ivan")
	fmt.Println(action.Human.Name) // Ivan
}
