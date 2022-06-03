package main

// func main() {

// 	/*
// 		Пульт дистанционного управления - 
// 		это наш Вызывающий(Invoker);
// 		ему будет передан объект command, 
// 		который можно использовать для выполения запросов.
// 	*/
// 	remoteControl := &simpleRemoteControl{}

// 	/*
// 	Теперь мы создаем Light объект,
// 	это будет Получатель(Receiver) запроса
// 	*/
// 	light := &light{}

// 	/*
// 	Создание команды и передача ей Получателя
// 	*/
// 	lightOnCommand := newLightOnCommand(light)

// 	/*
// 	Передача команды Вызывающему(Invoker)
// 	*/
// 	remoteControl.setCommand(lightOnCommand)

// 	/*
// 	Имитируем нажатие кнопки
// 	*/
// 	remoteControl.buttonWasPressed()

// 	garage := &garage{}

// 	garageDoorOpenCommand := newGarageDoorOpenCommand(garage)

// 	/*
// 	Передаем новую команду вызывающему
// 	*/
// 	remoteControl.setCommand(garageDoorOpenCommand)

// 	remoteControl.buttonWasPressed()
// }
