package main

import (
	"fmt"
	. "math"
	"strconv"
)

func InputStartData() (*StartData, error) {
	startData := new(StartData)
	var a string
	var err error
	s := "Ошибка: Некорректный формат входных данных\n"
	fmt.Printf("Введите вес мяча в кГ = ")
	fmt.Scan(&a)
	startData.weight, err = strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Printf(s)
		return nil, err
	}
	fmt.Printf("Введите сопротивление воздуха в кГ = ")
	fmt.Scan(&a)
	startData.air_resistance, err = strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Printf(s)
		return nil, err
	}
	fmt.Printf("Введите начальную скорость мяча в м/c = ")
	fmt.Scan(&a)
	startData.initial_speed, err = strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Printf(s)
		return nil, err
	}
	fmt.Printf("Введите точность вычислений (10^-x) = ")
	fmt.Scan(&a)
	startData.accuracy, err = strconv.Atoi(a)
	if err != nil {
		fmt.Printf(s)
		return nil, err
	}
	return startData, nil
}
func CalculateData(startData *StartData) Data {
	data := new(Data)
	w := startData.weight
	k := startData.air_resistance
	v := startData.initial_speed
	weight_SI := w * g / 1000
	air_resistance_SI := k * g
	mass := weight_SI / g * 1000
	data.mass = mass
	data.air_resistance_coefficient = air_resistance_SI
	data.initial_speed = v
	data.Air_resistance_asc()
	data.No_air_resistance()
	data.Air_resistance_desc()
	data.accuracy = startData.accuracy
	return *data
}
func (data *Data) PrintData(acc int) {
	fmt.Printf("----------С учётом сопротивления воздуха----------\n")
	fmt.Printf("Время подъёма: %s секунд\n", Format(data.time_resistance_asc, acc))
	fmt.Printf("Время падения: %s секунд\n", Format(data.time_resistance_desc, acc))
	fmt.Printf("Максимальная высота подъёма: %s метров\n", Format(data.height_resistance, acc))
	fmt.Printf("---------Без учёта сопротивления воздуха----------\n")
	fmt.Printf("Время подъёма: %s секунд\n", Format(data.time_no_resistance, acc))
	fmt.Printf("Время падения: %s секунд\n", Format(data.time_no_resistance, acc))
	fmt.Printf("Максимальная высота подъёма: %s метров\n", Format(data.height_no_resistance, acc))
}
func HeightResistanceAsc(data *Data, t float64) float64 { // Высота от времени с учетом сопротивления (подъём)
	m := data.mass
	k := data.air_resistance_coefficient
	v := data.initial_speed
	height := m / k *
		Log(
			Cos(
				Atan(v*(Sqrt(k/(m*g))))-
					t*(Sqrt(g*k/m)))/
				Cos(
					Atan(v*(Sqrt(k/(m*g))))))
	return height
}
func HeightResistanceDesc(data *Data, t float64) float64 { // Высота от времени с учетом сопротивления (спуск)
	t = t - data.time_resistance_asc
	m := data.mass
	k := data.air_resistance_coefficient
	h := data.height_resistance
	G := g * m / k
	km := k / m
	s := 2 * km * Sqrt(G)
	height := h + 2*Log(2)/(s)*Sqrt(G) - Sqrt(G)*(2*Log(Exp(s*t)+1)/s-t)
	return height
}
func HeightNoResistance(data *Data, t float64) float64 { // Высота от времени без учёта сопротивления
	v := data.initial_speed
	height := v*t - g*t*t/2
	return height
}
func (data *Data) Air_resistance_asc() {
	m := data.mass
	k := data.air_resistance_coefficient
	v := data.initial_speed
	data.time_resistance_asc = Sqrt(m/(g*k)) *
		Atan(v*(Sqrt(k/(m*g))))
	data.height_resistance = m / k *
		Log(
			Cos(
				Atan(v*(Sqrt(k/(m*g))))-
					data.time_resistance_asc*(Sqrt(g*k/m)))/
				Cos(
					Atan(v*(Sqrt(k/(m*g))))))
}
func (data *Data) Air_resistance_desc() {
	m := data.mass
	k := data.air_resistance_coefficient
	km := k / m
	h := data.height_resistance
	G := g * m / k
	V := Sqrt(G - G/(Exp(2*h*km)))
	data.time_resistance_desc = 1 / (2 * km * Sqrt(G)) * Log((Sqrt(G)+V)/(Sqrt(G)-V))
}
func (data *Data) No_air_resistance() Data {
	v := data.initial_speed
	t := v / g
	data.time_no_resistance = t
	data.height_no_resistance = v*t - g*t*t/2
	return *data
}
func Format(n float64, acc int) string { // Форматированный вывод с учётом точности вычислений
	format := "%." + fmt.Sprint(acc) + "f"
	return fmt.Sprintf(format, n)
}
