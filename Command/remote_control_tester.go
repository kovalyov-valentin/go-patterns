package main

import (
	"fmt"
	"reflect"
)

type iCommand interface {
	execute()
}

type lightOnCommand struct {
	light *light
}

/*
Конструктору передается конекретный источник света,
который будет управлять эта команда (к примеру свет в гостинной)
и прячет его в переменной экземпляра light.
Когда вызывается execute, это будет объект light, который будет получателем запроса.
*/
func newLightOnCommand(l *light) *lightOnCommand {
	return &lightOnCommand{
		light: l,
	}
}

/* 
Метод execute вызывает метод on() на принимающем объекте,
которым является light, которым мы управляем.
*/ 
func (l *lightOnCommand) execute() {
	l.light.on()
}

/* 
Команда lightOff работает точно также как и lightOn,
за тем лишь исключением, что привязка получателя происходит
к другому действию: методу off()
*/
type lightOffCommand struct {
	light *light 
}

func newLightOffCommand(l *light) *lightOffCommand {
	return &lightOffCommand{
		light: l,
	}
}

func (l *lightOffCommand) execute() {
	l.light.off()
}

type garageDoorOpenCommand struct {
	garage *garage 
}

func newGarageDoorOpenCommand(g *garage) *garageDoorOpenCommand {
	return &garageDoorOpenCommand{
		garage: g,
	}
}

func (g *garageDoorOpenCommand) execute() {
	g.garage.up()
	g.garage.lightOn()
}

type garageDoorCloseCommand struct {
	garage *garage 
}

func newGarageDoorCloseCommand(g *garage) *garageDoorCloseCommand {
	return &garageDoorCloseCommand{
		garage: g,
	}
}

func (g *garageDoorCloseCommand) execute() {
	g.garage.down()
	g.garage.lightOff()
}

/* 
Точно также, как и в команде lightOn, 
нам передается экземпляр stereo, которым мы собираемся управлять,
и мы сохраняем его в локальной переменной экземпляра
*/ 
type stereoOnCommand struct {
	stereo *stereo
}

func newStereoOnCommand(s *stereo) *stereoOnCommand {
	return &stereoOnCommand{
		stereo: s,
	}
}

/*
Чтобы выполнить этотт запрос, нам нужно вызвать три метода
на стереосистеме: сначала включить ее, настроить на
воспроизведение СД диска и установить на громкость 11
*/
func (s *stereoOnCommand) execute() {
	s.stereo.on()
	s.stereo.setCD()
	s.stereo.setVolume(11)
}

type stereoOffCommand struct {
	stereo *stereo
}

func newStereoOffCommand(s *stereo) *stereoOffCommand {
	return &stereoOffCommand{
		stereo: s,
	}
}

func (s *stereoOffCommand) execute() {
	s.stereo.off()
}

type ceilingFanOnCommand struct {
	ceilingFan *ceilingFan
}

func newCeilingFanOnCommand(c *ceilingFan) *ceilingFanOnCommand {
	return &ceilingFanOnCommand{
		ceilingFan: c,
	}
}

func (c *ceilingFanOnCommand) execute() {
	c.ceilingFan.on()
}

type ceilingFanOffCommand struct {
	ceilingFan *ceilingFan
}

func newCeilingFanOffCommand(c *ceilingFan) *ceilingFanOffCommand {
	return &ceilingFanOffCommand{
		ceilingFan: c,
	}
}

func (c *ceilingFanOffCommand) execute() {
	c.ceilingFan.off()
}

type macroCommand struct {
	commands []iCommand
}

func newMacroCommand(c []iCommand) *macroCommand {
	return &macroCommand{
		commands: c,
	}
}

func (m *macroCommand) addCommand(c iCommand) {
	m.commands = append(m.commands, c)
}

func (m *macroCommand) execute() {
	for _, command := range m.commands {
		command.execute()
	}
}

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

type simpleRemoteControl struct {
	slot iCommand
}
/*
Это метод для установки
команды, которой будет 
управлять слот. Это может быть вызвано
несколько раз, если клиент
этого кода захочет изменить
поведение удаленной кнопки.
*/
func (r *simpleRemoteControl) setCommand(command iCommand) {
	r.slot = command 
}

/*
Этот метод вызывается при
нажатии кнопки. Все, что мы делаем, это берем
текущую команду, привязанную к 
слоту, и весь ее метод execute()
*/
func (r *simpleRemoteControl) buttonWasPressed() {
	r.slot.execute()
}


func main() {
	remoteControl := newRemoteControl()

	livingRoomLight := &light{"Living Room"}
	kitchenLight := &light{"Kitchen"}
	ceilingFan := &ceilingFan{"Living Room"}
	garage := &garage{}
	stereo := &stereo{"Living Room"}


	livingRoomLightOn := newLightOnCommand(livingRoomLight)
	livingRoomLightOff := newLightOffCommand(livingRoomLight)
	kitchenLightOn := newLightOnCommand(kitchenLight)
	kitchenLightOff := newLightOffCommand(kitchenLight)


	ceilingFanOn := newCeilingFanOnCommand(ceilingFan)
	ceilingFanOff := newCeilingFanOffCommand(ceilingFan)

	stereoOnWithCD := newStereoOffCommand(stereo)
	stereoOff := newStereoOffCommand(stereo)

	garageDoorOpen := newGarageDoorOpenCommand(garage)
	garageDoorClose := newGarageDoorCloseCommand(garage)

	livingRoomOnCommands := []iCommand{livingRoomLightOn, ceilingFanOn, stereoOnWithCD}
	livingRoomOffCommands := []iCommand{livingRoomLightOff, ceilingFanOff, stereoOff}
	livingRoomMacroOnCommand := newMacroCommand(livingRoomOnCommands)
	livingRoomMacroOffCommand := newMacroCommand(livingRoomOffCommands)

	remoteControl.setCommand(0, livingRoomLightOn, livingRoomLightOff)
	remoteControl.setCommand(1, kitchenLightOn, kitchenLightOff)
	remoteControl.setCommand(2, ceilingFanOn, ceilingFanOff)
	remoteControl.setCommand(3, stereoOnWithCD, stereoOff)
	remoteControl.setCommand(4, garageDoorOpen, garageDoorClose)
	remoteControl.setCommand(5, livingRoomMacroOnCommand, livingRoomMacroOffCommand)

	fmt.Println(remoteControl)

	remoteControl.onButtonWasPushed(0)
	remoteControl.offButtonWasPushed(0)
	remoteControl.onButtonWasPushed(1)
	remoteControl.offButtonWasPushed(1)
	remoteControl.onButtonWasPushed(2)
	remoteControl.offButtonWasPushed(2)
	remoteControl.onButtonWasPushed(3)
	remoteControl.offButtonWasPushed(3)
	remoteControl.onButtonWasPushed(4)
	remoteControl.offButtonWasPushed(4)

	fmt.Println("---Pushing Macro On---")
	remoteControl.onButtonWasPushed(5)

	fmt.Println("---Pushing Macro Off---")
	remoteControl.offButtonWasPushed(5)
	
	// /*
	// 	Пульт дистанционного управления - 
	// 	это наш Вызывающий(Invoker);
	// 	ему будет передан объект command, 
	// 	который можно использовать для выполения запросов.
	// */
	// remoteControl := &simpleRemoteControl{}

	// /*
	// Теперь мы создаем Light объект,
	// это будет Получатель(Receiver) запроса
	// */
	// light := &light{}

	// /*
	// Создание команды и передача ей Получателя
	// */
	// lightOnCommand := newLightOnCommand(light)

	// /*
	// Передача команды Вызывающему(Invoker)
	// */
	// remoteControl.setCommand(lightOnCommand)

	// /*
	// Имитируем нажатие кнопки
	// */
	// remoteControl.buttonWasPressed()

	// garage := &garage{}

	// garageDoorOpenCommand := newGarageDoorOpenCommand(garage)

	// /*
	// Передаем новую команду вызывающему
	// */
	// remoteControl.setCommand(garageDoorOpenCommand)

	// remoteControl.buttonWasPressed()
}