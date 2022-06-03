package main

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