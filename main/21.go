package main

import "fmt"

func main() {
	weather := new(Weather)           //Представим, что структура Weather это неизменяемая библиотека
	weatherUtils := new(WeatherUtils) //Нам нужны данные о погоде, поделенной на 2, поэтому мы пишем структуру-обертку
	fmt.Println(weatherUtils.DivideWeather(weather))
}

type Weather struct {
}

func (weather *Weather) GetTodayData() float64 {
	return -25.7
}

type WeatherUtils struct {
}

func (weatherUtils *WeatherUtils) DivideWeather(weather *Weather) float64 {
	return weather.GetTodayData() / 2
}
