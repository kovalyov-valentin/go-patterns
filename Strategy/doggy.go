// package main

// import (
// 	"fmt"
// )

// type Dog struct {
// 	barker        Barker
// 	waggingTailer WaggingTailer
// 	sniffinger    Sniffinger
// }

// type IDog interface {
// 	Display()
// 	Barking()
// 	WaggingTail()
// 	Sniffing()
// }

// func (d *Dog) Display() {
// 	fmt.Println("Я собака")
// }

// func (d *Dog) Barking() {
// 	d.barker.bark()
// }

// func (d *Dog) WaggingTail() {
// 	d.waggingTailer.wagging()
// }

// func (d *Dog) Sniffing() {
// 	d.sniffinger.snif()
// }

// type Barker interface {
// 	bark()
// }

// type Bark struct {
// }

// func (b Bark) bark() {
// 	fmt.Println("Я лаю")
// }

// type WaggingTailer interface {
// 	wagging()
// }

// type WaggingTail struct {
// }

// func (w WaggingTail) wagging() {
// 	fmt.Println("Я виляю хвостом")
// }

// type Sniffinger interface {
// 	snif()
// }

// type Sniffing struct {
// }

// func (s Sniffing) snif() {
// 	fmt.Println("Я нюхаю жопы")
// }

// type Koba struct {
// 	Dog
// }

// func NewKoba() IDog {
// 	return &Koba{Dog{
// 		barker: Bark{}, 
// 		waggingTailer: WaggingTail{}, 
// 		sniffinger: Sniffing{}}}
// }

// func (k *Koba) Display() {
// 	fmt.Println("Пекинес")
// }

// func main() {
// 	koba := NewKoba()

// 	dogs := []IDog{koba}
// 	for _, dog := range dogs {
// 		fmt.Printf("\n%T\n", dog)
// 		dog.Display()
// 		dog.Barking()
// 		dog.WaggingTail()
// 		dog.Sniffing()
// 	}
// }
