package main 

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