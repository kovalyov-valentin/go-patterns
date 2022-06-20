package main

import (
	"fmt"
	"reflect"
)

/*
Пульт дистанционного управления будет обрабатывать
команды включения и выключения, которые будут
сохранены в соответствующих массивах
*/
type remoteControl struct {
	onCommands, offCommands []iCommand
}

func newRemoteControl() *remoteControl {
	return &remoteControl{
		onCommands:  make([]iCommand, 7),
		offCommands: make([]iCommand, 7),
	}
}

/*
Метод setCommand() принимает позицию слота
и команды включения и выключения, которые будут сохранены
в этом слоте.
Он помещает эти команды в массивы включения и выключения
для последующего использования
*/
func (r *remoteControl) setCommand(slot int, onCommands, offCommands iCommand) {
	r.onCommands[slot] = onCommands
	r.offCommands[slot] = offCommands
}

/*
Когда нажата кнопка включения или выключения,
аппаратное обеспечение заботится о вызове
соответсвующих методов при нажании button() или
при нажатии button() выключения
*/
func (r *remoteControl) onButtonWasPushed(slot int) {
	fmt.Println("*****")
	r.onCommands[slot].execute()
	fmt.Println("*****")
}

func (r *remoteControl) offButtonWasPushed(slot int) {
	fmt.Println("*****")
	r.offCommands[slot].execute()
	fmt.Println("*****")
}

/*
Реализация String() для распечатки каждого слота и
соответсвующей ему команды.
*/
func (r *remoteControl) String() string {
	s := fmt.Sprintf("\n------ Remote Control ------\n")

	for i := range r.onCommands {
		if r.onCommands[i] == nil {
			continue
		}

		onClass := r.getClassName(r.onCommands[i])
		offClass := r.getClassName(r.offCommands[i])
		s += fmt.Sprintf("[slot %d] %s   %s\n", i, onClass, offClass)
	}
	s += fmt.Sprintf("-----------------------------\n")

	return s
}

func (r *remoteControl) getClassName(myVar interface{}) string {
	if t := reflect.TypeOf(myVar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

