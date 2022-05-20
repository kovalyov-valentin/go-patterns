package main

import (
	"fmt"
)

type Subjject interface {
	Register(ob Observer)
	Remove(ob Observer)
	notifyAll()
}

type WeatherData struct {
	observers   []Observer
	temperature float32
	humidity    float32
	pressure    float32
}

func NewWeatherData() WeatherData {
	wd := WeatherData{}
	return wd 
}

func (w *WeatherData) Register(ob Observer) {
	w.observers = append(w.observers, ob)
}

func (w *WeatherData) Remove(ob Observer) {
	i := search(&w.observers, ob)
	w.observers = remove(w.observers, i)
}

func (w *WeatherData) notifyAll() {
	for _, observer := range w.observers {
		observer.update(w.temperature, w.humidity, w.pressure)
	}
}

func (w *WeatherData) measurementChanged() {
	w.notifyAll()
}

func (w *WeatherData) SetMeasurements(temperature float32, humidity float32, pressure float32) {
	w.temperature = temperature
	w.humidity = humidity 
	w.pressure = pressure 
	w.measurementChanged()
}
// Полезные функции (для чего полезные еще не понятно)
func remove(obs []Observer, i int) []Observer {
	obs[len(obs)-1], obs[i] = obs[i], obs[len(obs)-1]
	return obs[:len(obs)-1]
}

func search(obs *[]Observer, key Observer) int {
	for id, ob := range *obs {
		if ob == key {
			return id
		}
	}
	return -1
}

// Наблюдатель: общий интерфейс для всех наблюдателей
type Observer interface {
	update(temperature float32, humidity float32, pressure float32)
}

type Displayer interface {
	display()
}

type displayerObserver interface {
	Displayer 
	Observer 
}

// Отображение статистики
type StatDisplay struct {
	minTemp float32 
	maxTemp float32 
	tempSum float32 
	numReadings int 
	//weatherData WeatherData
}

func NewStatDisplay(wd WeatherData) Observer {
	sd := StatDisplay{minTemp: 200, maxTemp: 0.0, tempSum: 0.0}
	return &sd 
}

func (sd *StatDisplay) update(temp float32, humidity float32, pressure float32) {
	sd.tempSum = sd.tempSum + temp 
	sd.numReadings = sd.numReadings + 1

	if temp < sd.minTemp {
		sd.minTemp = temp 
	}

	if temp > sd.maxTemp {
		sd.maxTemp = temp
	}
	
	// Другой способ достижения множественного наследования
	var d Displayer = sd 
	d.display()
}

func (sd *StatDisplay) display() {
	fmt.Printf("Avg/Min/Max temperature = %0.1f/%0.1f/%0.1f\n", (sd.tempSum / float32(sd.numReadings)), sd.minTemp, sd.maxTemp)
}

// Отображение прогноза
type ForecastDisplay struct {
	currPressure float32 
	lastPressure float32 
	//weatherData WeatherData
}

func NewForecastDisplay(wd WeatherData) displayerObserver {
	fd := ForecastDisplay{}
	return &fd
}

func (fd *ForecastDisplay) update(temperature float32, humidity float32, pressure float32) {
	fd.lastPressure = fd.currPressure
	fd.currPressure = pressure 

	fd.display()
}

func (fd *ForecastDisplay) display() {
	fmt.Print("Forecast : ")
	if fd.currPressure > fd.lastPressure {
		fmt.Println("Improving weather on the way!")
	} else if fd.currPressure < fd.lastPressure {
		fmt.Println("Watch out for cooler, rainy weather!")
	} else {
		fmt.Println("More of the same.")
	}
}

// Отображение теплового индекса
type HeatIdxDisplay struct {
	heatIdx float32
	//weatherData WeatherData
}

func NewHeatIdxDisplay(wd WeatherData) Observer {
	hd := HeatIdxDisplay{}
	return &hd
}

func (hd *HeatIdxDisplay) display() {
	fmt.Printf("Heat index is %f\n", hd.heatIdx)
}

func (hd *HeatIdxDisplay) update(temperature float32, humidity float32, pressure float32) {
	hd.heatIdx = calculateHeatIndex(temperature, humidity)
	hd.display()
}

// Полезные функции
func calculateHeatIndex(t float32, rh float32) float32 {
	return ((16.923 + (0.185212 * t) + (5.37941 * rh) - (0.100254 * t * rh) +
	(0.00941695 * (t * t)) + (0.00728898 * (rh * rh)) +
	(0.000345372 * (t * t * rh)) - (0.000814971 * (t * rh * rh)) +
	(0.0000102102 * (t * t * rh * rh)) - (0.000038646 * (t * t * t)) + (0.0000291583 *
	(rh * rh * rh)) + (0.00000142721 * (t * t * t * rh)) +
	(0.000000197483 * (t * rh * rh * rh)) - (0.0000000218429 * (t * t * t * rh * rh)) +
	0.000000000843296*(t*t*rh*rh*rh)) -
	(0.0000000000481975 * (t * t * t * rh * rh * rh)))
}

type CurrCondDisplay struct {
	temperature float32
	humidity float32
}

func NewCurrCondDisplay(wd WeatherData) Observer {
	ccd := CurrCondDisplay{}
	return &ccd
}

func (ccd *CurrCondDisplay) update(temperature float32, humidity float32, pressure float32) {
	ccd.temperature = temperature
	ccd.humidity = humidity 
	ccd.display()
}

func (ccd *CurrCondDisplay) display() {
	fmt.Printf("Current conditions: %0.1fF degrees and &0.1f%% humidity.\n", ccd.temperature, ccd.humidity)
}

func main() {
	wd := NewWeatherData()

	ccd := NewCurrCondDisplay(wd)
	wd.Register(ccd)
	sd := NewStatDisplay(wd)
	wd.Register(sd)
	fd := NewForecastDisplay(wd)
	wd.Register(fd)
	hd := NewHeatIdxDisplay(wd)
	wd.Register(hd)

	wd.SetMeasurements(80, 65, 30.4)
	wd.SetMeasurements(82, 70, 29.2)

	wd.Remove(hd)

	wd.SetMeasurements(78, 90, 29.2)
}